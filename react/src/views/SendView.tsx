import React, { useEffect, useState } from "react";
import { useAccount, usePrepareSendTransaction, useSendTransaction, useWaitForTransaction } from "wagmi";
import { parseEther } from "ethers/lib/utils.js";
import { useDebounce } from "use-debounce";
import { Web3Button } from "@web3modal/react";
import { IgntButton } from "@ignt/react-library";
import { useClient } from "../hooks/useClient";
import { RegistryDomain, RegistryWalletRecordType } from "mycel-client-ts/mycel.registry/rest";
import { convertToNameAndParent } from "../utils/domainName";
import { useRegistryDomain } from "../def-hooks/useRegistryDomain";

const getWalletAddr = (domain: RegistryDomain, recordType: RegistryWalletRecordType) => {
  return domain?.walletRecords ? domain.walletRecords[recordType]?.value || "" : ""
}

export default function SendView() {
  const {registryDomain, isLoading: isLoadingRegistryDomain, updateRegistryDomain} = useRegistryDomain();
  const [domainName, setDomainName] = useState("")
  const [targetWalletRecordType, _] = useState(RegistryWalletRecordType.ETHEREUM_GOERLI)
  const [debouncedDomainName] = useDebounce(domainName, 500)
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

  const { isLoading: isLoadingTx, isSuccess } = useWaitForTransaction({
    hash: data?.hash,
  })

  useEffect(() => {
    if (registryDomain) {
      const walletAddr = registryDomain ? getWalletAddr(registryDomain, targetWalletRecordType) : ""
      setTo(walletAddr || "")
    } else {
      setTo("")
    }
  }, [registryDomain])

  useEffect(() => {
    updateRegistryDomain(domainName)
      .then(() => {})
      .catch(e => {
        console.error(e)
      })
  }, [debouncedDomainName])

  return (
    <div className="w-3/4 mx-auto">
      <h2 className=" text-2xl">Send GoerliETH DEMO</h2>
      <div className="m-4">
        <Web3Button />
      </div>
      <div className="flex-row m-4">
        <input
          className="mr-6 mt-2 py-2 px-4 h-14 bg-gray-100 w-full border-xs text-base leading-tight rounded-xl outline-0"
          aria-label="Recipient"
          onChange={async (e) => {
            setDomainName(e.target.value)
          }}
          placeholder="Recipient Mycel Domain Name(e.g. your-name.foo.cel)"
          value={domainName}
        />
        {
          to ? (
            <p className="m-2 text-sm text-gray-700"><span className="italic">{domainName}</span> in <span className="italic">{targetWalletRecordType}</span> is <span className="italic">{to}</span>.</p>
          ) : (
            <p className="m-2 text-sm text-red-500"><span className="italic">{domainName}</span> doesn't exists in registry.</p>
          )
        }
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
          busy={isLoadingTx || isLoadingRegistryDomain}
          disabled={isLoadingTx || isLoadingRegistryDomain || !sendTransactionAsync || !to || !amount}
        >
          {isLoadingTx ? 'Sending...' : 'Send'}
        </IgntButton>
        {isSuccess && (
          <div className="m-4">
            <p>Successfully sent {amount} ether to {to}</p>
            <div>
              <a className=" underline" href={`https://goerli.etherscan.io/tx/${data?.hash}`}>Etherscan Result Link</a>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
