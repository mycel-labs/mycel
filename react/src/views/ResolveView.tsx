import { useEffect, useState } from "react";
import { useClient } from "../hooks/useClient";
import { IgntButton } from "@ignt/react-library";
import { convertToDomainName, convertToNameAndParent } from "../utils/domainName";
import { useSearchParams } from 'react-router-dom';
import { useRegistryDomain } from "../def-hooks/useRegistryDomain";

export default function ResolveView() {
  const [query, setQuery] = useSearchParams({});
  const [inputtedDomainName, setInputtedDomainName] = useState("");
  const {registryDomain, isLoading, updateRegistryDomain} = useRegistryDomain();

  const updateRegistryHandler = async (domainName: string) => {
    try {
      await updateRegistryDomain(domainName)
      // Update query
      const { name, parent } = convertToNameAndParent(domainName);
      query.set("name", name)
      query.set("parent", parent)
      setQuery(query)
    } catch(e) {
      console.log(e)
      // Clear query
      query.delete("name")
      query.delete("parent")
      setQuery(query)
    }
  }

  useEffect(() => {
    const name = query.get("name") || ""
    const parent = query.get("parent") || ""
    if (inputtedDomainName || !name || !parent) {
      return
    }
    const domainName = convertToDomainName(name, parent)
    setInputtedDomainName(domainName)
    updateRegistryHandler(domainName)
      .then(() => {})
      .catch(e => {
        console.log(e)
      })
  }, [])

  return (
    <div className="w-3/4 mx-auto">
      {/* Uncomment the following component to add a form for a `modelName` -*/}
      {/* (<IgntCrud storeName="OrgRepoModule" itemName="modelName" />) */}
      <div className="flex mt-2 p-2 justify-between">
        <input
          className="mr-6 mt-1 py-2 px-4 h-14 bg-gray-100 w-full border-xs text-base leading-tight rounded-xl outline-0"
          placeholder="Mycel Domain"
          onChange={(event) => {
            setInputtedDomainName(event.target.value);
          }}
          onKeyDown={async (event) => {
            if (event.nativeEvent.isComposing || event.key !== 'Enter') return
            updateRegistryHandler(inputtedDomainName)
          }}
          value={inputtedDomainName}
        />
        <IgntButton className="mt-1 h-14 w-48"
          onClick={async () => {
            updateRegistryHandler(inputtedDomainName)
          }} busy={isLoading}>
          Resolve
        </IgntButton>
      </div>
      <div className="m-2">
        <div className="my-8">
          <h2 className=" text-2xl m-2 font-semibold">Basic Information</h2>
          <div className="table w-full border-collapse">
            <div className="table-header-group border-b font-medium">
              <div className=" table-cell w-4/12 p-2">Domain Name</div>
              <div className=" table-cell w-5/12 p-2">Owner Address</div>
              <div className=" table-cell w-3/12 p-2">Expiration Date</div>
            </div>
            <div className=" table-row">
              <div className="table-cell p-2">{convertToDomainName(registryDomain?.name, registryDomain?.parent)}</div>
              <div className="table-cell p-2">{registryDomain?.owner}</div>
              <div className="table-cell p-2">{registryDomain?.expirationDate ? (new Date(Math.round(parseInt(registryDomain?.expirationDate) / 1000000))).toUTCString() : ("")}</div>
            </div>
          </div>
        </div>
        <div className="my-8">
          <h2 className=" text-2xl m-2 font-semibold">DNS Records</h2>
          <div className="table w-full border-collapse">
            <div className="table-header-group border-b font-medium">
              <div className=" table-cell p-2">DNS Record Type</div>
              <div className=" table-cell p-2">Value</div>
            </div>
            {Object.values(registryDomain?.DNSRecords || []).map((v, i) => {
              return <div key={i} className="table-row text-justify">
                <div className="table-cell p-2">{v.DNSRecordType}</div>
                <div className="table-cell p-2">{v.value}</div>
              </div>
            })}
          </div>
        </div>
        <div className="my-8">
          <h2 className="text-2xl m-2 font-semibold">Wallet Records</h2>
          <div className="table w-full border-collapse">
            <div className="table-header-group border-b font-medium">
              <div className=" table-cell p-2">Wallet Record Type</div>
              <div className=" table-cell p-2">Wallet Address Format</div>
              <div className=" table-cell p-2">Value</div>
            </div>
            {Object.values(registryDomain?.walletRecords || []).map((v, i) => {
              return <div key={i} className=" table-row text-justify">
                <div className="table-cell p-2">{v.walletRecordType}</div>
                <div className="table-cell p-2">{v.WalletAddressFormat}</div>
                <div className="table-cell p-2">{v.value}</div>
              </div>
            })}
          </div>
        </div>
      </div>
    </div>
  );
}
