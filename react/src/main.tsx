import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";
import router from "./router";
import "./index.css";
import "@ignt/react-library/dist/style.css";
import AddressProvider from "./def-hooks/addressContext";

import DenomProvider from "./def-hooks/denomContext";
import WalletProvider from "./def-hooks/walletContext";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import { EthereumClient, w3mConnectors, w3mProvider } from '@web3modal/ethereum'
import { Web3Modal } from '@web3modal/react'
import { configureChains, createClient, WagmiConfig } from 'wagmi'
import { mainnet, polygon, goerli, polygonMumbai, gnosisChiado } from 'wagmi/chains'

const queryClient = new QueryClient();

const chains = [mainnet, polygon, goerli, polygonMumbai, gnosisChiado]
const projectId = '16f0505118c2c30cff533282efbde193'

const { provider } = configureChains(chains, [w3mProvider({ projectId })])
const wagmiClient = createClient({
  autoConnect: true,
  connectors: w3mConnectors({ projectId, version: 1, chains }),
  provider
})
const ethereumClient = new EthereumClient(wagmiClient, chains)

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <WagmiConfig client={wagmiClient}>
      <QueryClientProvider client={queryClient}>
        <AddressProvider>
          <WalletProvider>
            <DenomProvider>
              <RouterProvider router={router} />
            </DenomProvider>
          </WalletProvider>
        </AddressProvider>
      </QueryClientProvider>
    </WagmiConfig>
    <Web3Modal projectId={projectId} ethereumClient={ethereumClient} />
  </React.StrictMode>,
);
