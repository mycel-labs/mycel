import { useEffect, useState } from "react";
import { useClient } from "../hooks/useClient";
import { RegistryDomain } from "mycel-client-ts/mycel.registry/rest";
import { IgntButton } from "@ignt/react-library";
import { getNameAndParent } from "../utils/getNameAndParent";
import { useLocation } from 'react-router-dom';

export default function ResolveView() {
  const search = useLocation().search;
  const query = new URLSearchParams(search);
  const client = useClient();
  const [isLoading, setIsLoading] = useState(false);
  const [inputName, setInputName] = useState("");
  const [name, setName] = useState("");
  const [parent, setParent] = useState("");
  const [domain, setDomain] = useState<RegistryDomain | null>(null)

  const getDomain = async (name: string, parent: string) => {
    setIsLoading(true);
    try {
      const domain = await client.MycelRegistry.query.queryDomain(name, parent);
      setDomain(domain.data.domain || null);
      console.log(domain.data.domain);
    } catch (e) {
      console.error(e);
      setDomain(null);
      setIsLoading(false);
      return;
    }
    setIsLoading(false);
  }

  useEffect(() => {
    const name = query.get("name") || ""
    const parent = query.get("parent") || ""
    if (inputName !== "" || name === "" || parent === "") {
      return
    }

    setInputName(name + "." + parent)
    setName(name)
    setParent(parent)

    getDomain(name, parent)
      .then(() => { })
      .catch(e => { })
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
            const { name, parent } = getNameAndParent(event.target.value);
            setInputName(event.target.value);
            setName(name);
            setParent(parent);
          }}
          onKeyDown={async (event) => {
            if (event.nativeEvent.isComposing || event.key !== 'Enter') return
            await getDomain(name, parent);
            console.log("enter!")
          }}
          value={inputName}
        />
        <IgntButton className="mt-1 h-14 w-48"
          onClick={async () => { await getDomain(name, parent) }} busy={isLoading}>
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
              <div className="table-cell p-2">{domain?.name}.{domain?.parent}</div>
              <div className="table-cell p-2">{domain?.owner}</div>
              <div className="table-cell p-2">{domain?.expirationDate ? (new Date(Math.round(parseInt(domain?.expirationDate) / 1000000))).toUTCString() : ("")}</div>
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
            {Object.values(domain?.DNSRecords || []).map((v, i) => {
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
            {Object.values(domain?.walletRecords || []).map((v, i) => {
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
