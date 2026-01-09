require("@nomicfoundation/hardhat-toolbox");
require('dotenv').config();

// Configure networks with optional testnet/mainnet
const networks = {
  hardhat: {},
  localhost: {
    url: "http://127.0.0.1:8545"
  }
};

// Add Sepolia only if configured
if (process.env.SEPOLIA_RPC_URL && process.env.PRIVATE_KEY) {
  networks.sepolia = {
    url: process.env.SEPOLIA_RPC_URL,
    accounts: [process.env.PRIVATE_KEY]
  };
}

// Add Mainnet only if configured
if (process.env.MAINNET_RPC_URL && process.env.PRIVATE_KEY) {
  networks.mainnet = {
    url: process.env.MAINNET_RPC_URL,
    accounts: [process.env.PRIVATE_KEY]
  };
}

module.exports = {
  solidity: {
    version: "0.8.20",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  networks: networks,
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  }
};