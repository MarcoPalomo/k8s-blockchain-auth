// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title K8sAccessControl
 * @dev Smart contract pour gérer les permissions Kubernetes via blockchain
 * Permet d'attribuer des permissions granulaires (namespaces, verbs, resources) à des wallets
 */
contract K8sAccessControl is AccessControl {
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    // Roles
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant PERMISSION_MANAGER_ROLE = keccak256("PERMISSION_MANAGER_ROLE");

    // Structure pour stocker les permissions d'un utilisateur
    struct Permission {
        string[] namespaces;    // Liste des namespaces (ex: ["default", "production"])
        string[] verbs;         // Actions K8s (ex: ["get", "list", "create", "delete"])
        string[] resources;     // Resources K8s (ex: ["pods", "services", "deployments"])
        uint256 expiresAt;      // Timestamp d'expiration (0 = jamais)
        bool isActive;          // Permission active ou révoquée
    }

    // Mapping wallet address => Permission
    mapping(address => Permission) private permissions;

    // Set de tous les wallets ayant des permissions
    EnumerableSet.AddressSet private authorizedWallets;

    // Events
    event PermissionGranted(
        address indexed wallet,
        string[] namespaces,
        string[] verbs,
        string[] resources,
        uint256 expiresAt
    );

    event PermissionRevoked(address indexed wallet);

    event PermissionUpdated(
        address indexed wallet,
        string[] namespaces,
        string[] verbs,
        string[] resources
    );

    constructor() {
        // Le déployeur devient l'admin principal
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
        _grantRole(PERMISSION_MANAGER_ROLE, msg.sender);
    }

    /**
     * @dev Attribuer des permissions à un wallet
     * @param wallet Adresse du wallet
     * @param namespaces Liste des namespaces autorisés (["*"] pour tous)
     * @param verbs Liste des verbs autorisés (["*"] pour tous)
     * @param resources Liste des resources autorisées (["*"] pour toutes)
     * @param expiresAt Timestamp d'expiration (0 pour jamais)
     */
    function grantPermission(
        address wallet,
        string[] memory namespaces,
        string[] memory verbs,
        string[] memory resources,
        uint256 expiresAt
    ) external onlyRole(PERMISSION_MANAGER_ROLE) {
        require(wallet != address(0), "Invalid wallet address");
        require(namespaces.length > 0, "At least one namespace required");
        require(verbs.length > 0, "At least one verb required");
        require(resources.length > 0, "At least one resource required");

        if (expiresAt > 0) {
            require(expiresAt > block.timestamp, "Expiration must be in the future");
        }

        permissions[wallet] = Permission({
            namespaces: namespaces,
            verbs: verbs,
            resources: resources,
            expiresAt: expiresAt,
            isActive: true
        });

        authorizedWallets.add(wallet);

        emit PermissionGranted(wallet, namespaces, verbs, resources, expiresAt);
    }

    /**
     * @dev Révoquer les permissions d'un wallet
     * @param wallet Adresse du wallet
     */
    function revokePermission(address wallet) external onlyRole(PERMISSION_MANAGER_ROLE) {
        require(permissions[wallet].isActive, "Wallet has no active permissions");

        permissions[wallet].isActive = false;
        authorizedWallets.remove(wallet);

        emit PermissionRevoked(wallet);
    }

    /**
     * @dev Mettre à jour les permissions d'un wallet existant
     * @param wallet Adresse du wallet
     * @param namespaces Nouvelle liste de namespaces
     * @param verbs Nouvelle liste de verbs
     * @param resources Nouvelle liste de resources
     */
    function updatePermission(
        address wallet,
        string[] memory namespaces,
        string[] memory verbs,
        string[] memory resources
    ) external onlyRole(PERMISSION_MANAGER_ROLE) {
        require(permissions[wallet].isActive, "Wallet has no active permissions");
        require(namespaces.length > 0, "At least one namespace required");
        require(verbs.length > 0, "At least one verb required");
        require(resources.length > 0, "At least one resource required");

        permissions[wallet].namespaces = namespaces;
        permissions[wallet].verbs = verbs;
        permissions[wallet].resources = resources;

        emit PermissionUpdated(wallet, namespaces, verbs, resources);
    }

    /**
     * @dev Vérifier si un wallet a une permission spécifique
     * @param wallet Adresse du wallet
     * @param namespace Namespace K8s
     * @param verb Action K8s (get, list, create, etc.)
     * @param resource Resource K8s (pods, services, etc.)
     * @return bool True si la permission est accordée
     */
    function hasPermission(
        address wallet,
        string memory namespace,
        string memory verb,
        string memory resource
    ) external view returns (bool) {
        Permission memory perm = permissions[wallet];

        // Vérifier si la permission est active
        if (!perm.isActive) {
            return false;
        }

        // Vérifier l'expiration
        if (perm.expiresAt > 0 && block.timestamp > perm.expiresAt) {
            return false;
        }

        // Vérifier le namespace
        if (!_arrayContains(perm.namespaces, namespace) &&
            !_arrayContains(perm.namespaces, "*")) {
            return false;
        }

        // Vérifier le verb
        if (!_arrayContains(perm.verbs, verb) &&
            !_arrayContains(perm.verbs, "*")) {
            return false;
        }

        // Vérifier la resource
        if (!_arrayContains(perm.resources, resource) &&
            !_arrayContains(perm.resources, "*")) {
            return false;
        }

        return true;
    }

    /**
     * @dev Récupérer toutes les permissions d'un wallet
     * @param wallet Adresse du wallet
     * @return namespaces Liste des namespaces
     * @return verbs Liste des verbs
     * @return resources Liste des resources
     * @return expiresAt Timestamp d'expiration
     */
    function getPermissions(address wallet)
        external
        view
        returns (
            string[] memory namespaces,
            string[] memory verbs,
            string[] memory resources,
            uint256 expiresAt
        )
    {
        Permission memory perm = permissions[wallet];

        if (!perm.isActive) {
            return (new string[](0), new string[](0), new string[](0), 0);
        }

        return (perm.namespaces, perm.verbs, perm.resources, perm.expiresAt);
    }

    /**
     * @dev Vérifier si un wallet a des permissions actives
     * @param wallet Adresse du wallet
     * @return bool True si le wallet a des permissions actives
     */
    function isAuthorized(address wallet) external view returns (bool) {
        Permission memory perm = permissions[wallet];

        if (!perm.isActive) {
            return false;
        }

        if (perm.expiresAt > 0 && block.timestamp > perm.expiresAt) {
            return false;
        }

        return true;
    }

    /**
     * @dev Obtenir le nombre total de wallets autorisés
     * @return uint256 Nombre de wallets
     */
    function getAuthorizedWalletsCount() external view returns (uint256) {
        return authorizedWallets.length();
    }

    /**
     * @dev Obtenir un wallet autorisé par son index
     * @param index Index du wallet
     * @return address Adresse du wallet
     */
    function getAuthorizedWalletAt(uint256 index) external view returns (address) {
        require(index < authorizedWallets.length(), "Index out of bounds");
        return authorizedWallets.at(index);
    }

    /**
     * @dev Vérifier si un tableau de strings contient une valeur
     * @param array Tableau de strings
     * @param value Valeur à rechercher
     * @return bool True si la valeur est trouvée
     */
    function _arrayContains(string[] memory array, string memory value)
        private
        pure
        returns (bool)
    {
        for (uint256 i = 0; i < array.length; i++) {
            if (keccak256(bytes(array[i])) == keccak256(bytes(value))) {
                return true;
            }
        }
        return false;
    }

    /**
     * @dev Attribuer le rôle de gestionnaire de permissions
     * @param account Adresse à qui attribuer le rôle
     */
    function addPermissionManager(address account) external onlyRole(ADMIN_ROLE) {
        grantRole(PERMISSION_MANAGER_ROLE, account);
    }

    /**
     * @dev Révoquer le rôle de gestionnaire de permissions
     * @param account Adresse à qui révoquer le rôle
     */
    function removePermissionManager(address account) external onlyRole(ADMIN_ROLE) {
        revokeRole(PERMISSION_MANAGER_ROLE, account);
    }
}
