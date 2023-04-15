import React, { useEffect, useState } from "react";
import { useAccount, usePrepareSendTransaction, useSendTransaction, useWaitForTransaction } from "wagmi";
import { parseEther } from "ethers/lib/utils.js";
import { useDebounce } from "use-debounce";
import { Web3Button } from "@web3modal/react";
import { IgntButton } from "@ignt/react-library";
import { useClient } from "../hooks/useClient";
import { RegistryDomain } from "mycel-client-ts/mycel.registry/rest";
import { getNameAndParent } from "../utils/getNameAndParent";

export default function SendView() {
  const client = useClient();
  const [domainName, setDomainName] = useState("")
  const [debouncedDomainName] = useDebounce(domainName, 300)
  const [to, setTo] = useState("")

  const [amount, setAmount] = useState("")
  const [debouncedAmount] = useDebounce(amount, 500)

  const { config } = usePrepareSendTransaction({
    request: {
      to: to,
      value: debouncedAmount ? parseEther(debouncedAmount) : undefined,
    },
  })
  const { data, sendTransactionAsync } = useSendTransaction(config)

  const { isLoading, isSuccess } = useWaitForTransaction({
    hash: data?.hash,
  })

  const getDomain = async (name: string, parent: string): Promise<RegistryDomain | null> => {
    try {
      const domain = await client.MycelRegistry.query.queryDomain(name, parent);
      return domain.data.domain || null
    } catch (e) {
      console.error(e);
      return null;
    }
  }

  const convertToWalletAddr = async () => {
    const {name, parent} = getNameAndParent(domainName)
    const domain = await getDomain(name, parent)
    const walletAddr = domain?.walletRecords ? domain.walletRecords["ETHEREUM_MAINNET"].value || "" : ""
    setTo(walletAddr || "")
  }

  useEffect(() => {
    convertToWalletAddr()
      .then(() => {})
      .catch(e => {console.log(e)})
  }, [debouncedDomainName])

  return (
    <div className="w-3/4 mx-auto">
      <div className="m-4">
        <Web3Button />
      </div>
      <div className="flex-row m-4">
        <input
          className="mr-6 my-2 py-2 px-4 h-14 bg-gray-100 w-full border-xs text-base leading-tight rounded-xl outline-0"
          aria-label="Recipient"
          onChange={async (e) => {
            setDomainName(e.target.value)
          }}
          placeholder="Recipient Mycel Domain Name(e.g. your-name.foo.cel)"
          value={domainName}
        />
        <input
          className="mr-6 my-2 py-2 px-4 h-14 bg-gray-100 w-full border-xs text-base leading-tight rounded-xl outline-0"
          aria-label="Amount (ether)"
          onChange={(e) => setAmount(e.target.value)}
          placeholder="Token Amount(e.g. 0.05)"
          value={amount}
        />

        <IgntButton className="mt-1 h-14 w-full"
          onClick={async () => {
            const res = await sendTransactionAsync?.()
            console.log("%o", res)
          }}
          busy={isLoading}
          disabled={isLoading || !sendTransactionAsync || !to || !amount}
        >
          {isLoading ? 'Sending...' : 'Send'}
        </IgntButton>
        {isSuccess && (
          <div className="m-4">
            Successfully sent {amount} ether to {to}
            <div>
              <a href={`https://goerli.etherscan.io/tx/${data?.hash}`}>Etherscan</a>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
