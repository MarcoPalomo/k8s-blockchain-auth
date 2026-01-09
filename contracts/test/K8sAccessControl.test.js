const { expect } = require("chai");
const { ethers } = require("hardhat");
const { time } = require("@nomicfoundation/hardhat-network-helpers");

describe("K8sAccessControl", function () {
  let K8sAccessControl;
  let contract;
  let owner;
  let permissionManager;
  let user1;
  let user2;

  beforeEach(async function () {
    // Récupérer les signataires
    [owner, permissionManager, user1, user2] = await ethers.getSigners();

    // Déployer le contrat
    K8sAccessControl = await ethers.getContractFactory("K8sAccessControl");
    contract = await K8sAccessControl.deploy();
    await contract.waitForDeployment();
  });

  describe("Deployment", function () {
    it("Should set the deployer as admin", async function () {
      const ADMIN_ROLE = await contract.ADMIN_ROLE();
      expect(await contract.hasRole(ADMIN_ROLE, owner.address)).to.be.true;
    });

    it("Should set the deployer as permission manager", async function () {
      const PERMISSION_MANAGER_ROLE = await contract.PERMISSION_MANAGER_ROLE();
      expect(await contract.hasRole(PERMISSION_MANAGER_ROLE, owner.address)).to.be.true;
    });
  });

  describe("Role Management", function () {
    it("Should allow admin to add permission manager", async function () {
      await contract.addPermissionManager(permissionManager.address);

      const PERMISSION_MANAGER_ROLE = await contract.PERMISSION_MANAGER_ROLE();
      expect(await contract.hasRole(PERMISSION_MANAGER_ROLE, permissionManager.address)).to.be.true;
    });

    it("Should allow admin to remove permission manager", async function () {
      await contract.addPermissionManager(permissionManager.address);
      await contract.removePermissionManager(permissionManager.address);

      const PERMISSION_MANAGER_ROLE = await contract.PERMISSION_MANAGER_ROLE();
      expect(await contract.hasRole(PERMISSION_MANAGER_ROLE, permissionManager.address)).to.be.false;
    });

    it("Should not allow non-admin to add permission manager", async function () {
      await expect(
        contract.connect(user1).addPermissionManager(permissionManager.address)
      ).to.be.reverted;
    });
  });

  describe("Grant Permission", function () {
    it("Should grant permissions to a wallet", async function () {
      const namespaces = ["default", "production"];
      const verbs = ["get", "list", "create"];
      const resources = ["pods", "services"];

      await expect(
        contract.grantPermission(user1.address, namespaces, verbs, resources, 0)
      )
        .to.emit(contract, "PermissionGranted")
        .withArgs(user1.address, namespaces, verbs, resources, 0);

      const [ns, vb, res, exp] = await contract.getPermissions(user1.address);
      expect(ns).to.deep.equal(namespaces);
      expect(vb).to.deep.equal(verbs);
      expect(res).to.deep.equal(resources);
      expect(exp).to.equal(0);
    });

    it("Should grant wildcard permissions", async function () {
      await contract.grantPermission(user1.address, ["*"], ["*"], ["*"], 0);

      expect(await contract.hasPermission(user1.address, "default", "get", "pods")).to.be.true;
      expect(await contract.hasPermission(user1.address, "production", "delete", "services")).to.be.true;
      expect(await contract.hasPermission(user1.address, "any-namespace", "any-verb", "any-resource")).to.be.true;
    });

    it("Should grant permissions with expiration", async function () {
      const futureTime = (await time.latest()) + 3600; // +1 hour
      await contract.grantPermission(
        user1.address,
        ["default"],
        ["get"],
        ["pods"],
        futureTime
      );

      const [, , , exp] = await contract.getPermissions(user1.address);
      expect(exp).to.equal(futureTime);
    });

    it("Should reject invalid wallet address", async function () {
      await expect(
        contract.grantPermission(ethers.ZeroAddress, ["default"], ["get"], ["pods"], 0)
      ).to.be.revertedWith("Invalid wallet address");
    });

    it("Should reject empty namespaces", async function () {
      await expect(
        contract.grantPermission(user1.address, [], ["get"], ["pods"], 0)
      ).to.be.revertedWith("At least one namespace required");
    });

    it("Should reject past expiration time", async function () {
      const pastTime = (await time.latest()) - 3600; // -1 hour
      await expect(
        contract.grantPermission(user1.address, ["default"], ["get"], ["pods"], pastTime)
      ).to.be.revertedWith("Expiration must be in the future");
    });

    it("Should not allow non-permission-manager to grant permissions", async function () {
      await expect(
        contract.connect(user1).grantPermission(user2.address, ["default"], ["get"], ["pods"], 0)
      ).to.be.reverted;
    });
  });

  describe("Permission Checking", function () {
    beforeEach(async function () {
      // Donner des permissions limitées à user1
      await contract.grantPermission(
        user1.address,
        ["default", "production"],
        ["get", "list"],
        ["pods", "services"],
        0
      );
    });

    it("Should correctly check valid permissions", async function () {
      expect(await contract.hasPermission(user1.address, "default", "get", "pods")).to.be.true;
      expect(await contract.hasPermission(user1.address, "production", "list", "services")).to.be.true;
    });

    it("Should reject invalid namespace", async function () {
      expect(await contract.hasPermission(user1.address, "kube-system", "get", "pods")).to.be.false;
    });

    it("Should reject invalid verb", async function () {
      expect(await contract.hasPermission(user1.address, "default", "create", "pods")).to.be.false;
    });

    it("Should reject invalid resource", async function () {
      expect(await contract.hasPermission(user1.address, "default", "get", "deployments")).to.be.false;
    });

    it("Should handle wildcard namespace", async function () {
      await contract.grantPermission(user2.address, ["*"], ["get"], ["pods"], 0);

      expect(await contract.hasPermission(user2.address, "default", "get", "pods")).to.be.true;
      expect(await contract.hasPermission(user2.address, "any-namespace", "get", "pods")).to.be.true;
    });

    it("Should handle wildcard verbs", async function () {
      await contract.grantPermission(user2.address, ["default"], ["*"], ["pods"], 0);

      expect(await contract.hasPermission(user2.address, "default", "get", "pods")).to.be.true;
      expect(await contract.hasPermission(user2.address, "default", "delete", "pods")).to.be.true;
    });

    it("Should handle wildcard resources", async function () {
      await contract.grantPermission(user2.address, ["default"], ["get"], ["*"], 0);

      expect(await contract.hasPermission(user2.address, "default", "get", "pods")).to.be.true;
      expect(await contract.hasPermission(user2.address, "default", "get", "services")).to.be.true;
    });
  });

  describe("Permission Expiration", function () {
    it("Should deny access after expiration", async function () {
      const futureTime = (await time.latest()) + 3600; // +1 hour
      await contract.grantPermission(user1.address, ["default"], ["get"], ["pods"], futureTime);

      // Vérifier que c'est accessible avant expiration
      expect(await contract.hasPermission(user1.address, "default", "get", "pods")).to.be.true;

      // Avancer le temps au-delà de l'expiration
      await time.increaseTo(futureTime + 1);

      // Vérifier que c'est maintenant refusé
      expect(await contract.hasPermission(user1.address, "default", "get", "pods")).to.be.false;
      expect(await contract.isAuthorized(user1.address)).to.be.false;
    });

    it("Should allow access with zero expiration (never expires)", async function () {
      await contract.grantPermission(user1.address, ["default"], ["get"], ["pods"], 0);

      // Avancer le temps de beaucoup
      await time.increase(365 * 24 * 60 * 60); // +1 year

      // Devrait toujours être accessible
      expect(await contract.hasPermission(user1.address, "default", "get", "pods")).to.be.true;
      expect(await contract.isAuthorized(user1.address)).to.be.true;
    });
  });

  describe("Update Permission", function () {
    beforeEach(async function () {
      await contract.grantPermission(user1.address, ["default"], ["get"], ["pods"], 0);
    });

    it("Should update existing permissions", async function () {
      const newNamespaces = ["production", "staging"];
      const newVerbs = ["get", "list", "create"];
      const newResources = ["services", "deployments"];

      await expect(
        contract.updatePermission(user1.address, newNamespaces, newVerbs, newResources)
      )
        .to.emit(contract, "PermissionUpdated")
        .withArgs(user1.address, newNamespaces, newVerbs, newResources);

      const [ns, vb, res] = await contract.getPermissions(user1.address);
      expect(ns).to.deep.equal(newNamespaces);
      expect(vb).to.deep.equal(newVerbs);
      expect(res).to.deep.equal(newResources);
    });

    it("Should not update non-existent permissions", async function () {
      await expect(
        contract.updatePermission(user2.address, ["default"], ["get"], ["pods"])
      ).to.be.revertedWith("Wallet has no active permissions");
    });
  });

  describe("Revoke Permission", function () {
    beforeEach(async function () {
      await contract.grantPermission(user1.address, ["default"], ["get"], ["pods"], 0);
    });

    it("Should revoke permissions", async function () {
      await expect(contract.revokePermission(user1.address))
        .to.emit(contract, "PermissionRevoked")
        .withArgs(user1.address);

      expect(await contract.isAuthorized(user1.address)).to.be.false;
      expect(await contract.hasPermission(user1.address, "default", "get", "pods")).to.be.false;
    });

    it("Should not revoke already revoked permissions", async function () {
      await contract.revokePermission(user1.address);

      await expect(
        contract.revokePermission(user1.address)
      ).to.be.revertedWith("Wallet has no active permissions");
    });

    it("Should remove wallet from authorized list", async function () {
      const countBefore = await contract.getAuthorizedWalletsCount();
      await contract.revokePermission(user1.address);
      const countAfter = await contract.getAuthorizedWalletsCount();

      expect(countAfter).to.equal(countBefore - 1n);
    });
  });

  describe("Authorized Wallets", function () {
    it("Should track authorized wallets count", async function () {
      expect(await contract.getAuthorizedWalletsCount()).to.equal(0);

      await contract.grantPermission(user1.address, ["default"], ["get"], ["pods"], 0);
      expect(await contract.getAuthorizedWalletsCount()).to.equal(1);

      await contract.grantPermission(user2.address, ["default"], ["get"], ["pods"], 0);
      expect(await contract.getAuthorizedWalletsCount()).to.equal(2);
    });

    it("Should get authorized wallet by index", async function () {
      await contract.grantPermission(user1.address, ["default"], ["get"], ["pods"], 0);
      await contract.grantPermission(user2.address, ["default"], ["get"], ["pods"], 0);

      const wallet0 = await contract.getAuthorizedWalletAt(0);
      const wallet1 = await contract.getAuthorizedWalletAt(1);

      expect([wallet0, wallet1]).to.include(user1.address);
      expect([wallet0, wallet1]).to.include(user2.address);
    });

    it("Should revert when index is out of bounds", async function () {
      await expect(
        contract.getAuthorizedWalletAt(0)
      ).to.be.revertedWith("Index out of bounds");
    });
  });

  describe("Complex Scenarios", function () {
    it("Should handle multiple users with different permissions", async function () {
      // Admin user
      await contract.grantPermission(user1.address, ["*"], ["*"], ["*"], 0);

      // Developer user
      await contract.grantPermission(
        user2.address,
        ["default", "development"],
        ["get", "list", "create", "update"],
        ["pods", "services"],
        0
      );

      // Vérifier admin
      expect(await contract.hasPermission(user1.address, "production", "delete", "deployments")).to.be.true;

      // Vérifier developer
      expect(await contract.hasPermission(user2.address, "default", "get", "pods")).to.be.true;
      expect(await contract.hasPermission(user2.address, "production", "get", "pods")).to.be.false;
      expect(await contract.hasPermission(user2.address, "default", "delete", "pods")).to.be.false;
    });

    it("Should handle permission lifecycle", async function () {
      // Grant
      await contract.grantPermission(user1.address, ["default"], ["get"], ["pods"], 0);
      expect(await contract.isAuthorized(user1.address)).to.be.true;

      // Update
      await contract.updatePermission(user1.address, ["production"], ["list"], ["services"]);
      expect(await contract.hasPermission(user1.address, "production", "list", "services")).to.be.true;
      expect(await contract.hasPermission(user1.address, "default", "get", "pods")).to.be.false;

      // Revoke
      await contract.revokePermission(user1.address);
      expect(await contract.isAuthorized(user1.address)).to.be.false;
    });
  });
});
