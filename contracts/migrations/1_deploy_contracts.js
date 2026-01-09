const hre = require("hardhat");

async function main() {
  console.log("Déploiement du contrat K8sAccessControl...");

  const [deployer] = await hre.ethers.getSigners();
  console.log("Déploiement avec le compte:", deployer.address);

  const balance = await hre.ethers.provider.getBalance(deployer.address);
  console.log("Balance du compte:", hre.ethers.formatEther(balance), "ETH");

  const K8sAccessControl = await hre.ethers.getContractFactory("K8sAccessControl");
  const contract = await K8sAccessControl.deploy();

  await contract.waitForDeployment();

  const contractAddress = await contract.getAddress();
  console.log("Contrat K8sAccessControl déployé à l'adresse:", contractAddress);

  console.log("\n=== Configuration pour votre API Server ===");
  console.log(`export CONTRACT_ADDRESS="${contractAddress}"`);
  console.log(`export BLOCKCHAIN_RPC_URL="${hre.network.config.url}"`);

  const fs = require("fs");
  const path = require("path");

  const deploymentInfo = {
    network: hre.network.name,
    contractAddress: contractAddress,
    deployer: deployer.address,
    timestamp: new Date().toISOString(),
    blockNumber: await hre.ethers.provider.getBlockNumber()
  };

  const deploymentsDir = path.join(__dirname, "../../deployments");
  if (!fs.existsSync(deploymentsDir)) {
    fs.mkdirSync(deploymentsDir, { recursive: true });
  }

  const deploymentFile = path.join(deploymentsDir, `${hre.network.name}.json`);
  fs.writeFileSync(deploymentFile, JSON.stringify(deploymentInfo, null, 2));
  console.log(`\nInformations de déploiement sauvegardées dans: ${deploymentFile}`);

  const fixturesDir = path.join(__dirname, "../../test/fixtures/contracts");
  if (!fs.existsSync(fixturesDir)) {
    fs.mkdirSync(fixturesDir, { recursive: true });
  }
  const fixturesFile = path.join(fixturesDir, "deployed-addresses.json");
  fs.writeFileSync(fixturesFile, JSON.stringify(deploymentInfo, null, 2));

  console.log("\n✅ Déploiement terminé avec succès!");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
