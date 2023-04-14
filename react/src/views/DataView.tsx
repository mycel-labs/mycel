import { useEffect, useState } from "react";
import IgntCrud from "../components/IgntCrud";
import useMycelRegistry from "../hooks/useMycelRegistry";
import { useRegistryDomain } from "../def-hooks/useRegistryDomain";
import { useClient } from "../hooks/useClient";
import { RegistryDomain } from "mycel-client-ts/mycel.registry/rest";

const getNameAndParent = (domain: string) => {
  const s = domain.split(".");
  if (s.length === 1) {
    return { name: "", parent: s[0] };
  }
  return { name: s[0], parent: s.slice(1).join(".") }
}

export default function DataView() {
  const client = useClient();
  const [name, setName] = useState("");
  const [parent, setParent] = useState("");
  const [domain, setDomain] = useState<RegistryDomain | null>(null)

  const getDomain = async (name: string, parent: string) => {
    const domain = await client.MycelRegistry.query.queryDomain(name, parent);

    setDomain(domain.data.domain || null);
  }

  useEffect(() => {
    getDomain(name, parent)
      .then()
      .catch(e => {
        console.log(e)
        setDomain(null)
      })
    console.log(name, parent)
  }, [name, parent])

  return (
    <div className="m-2">
      {/* Uncomment the following component to add a form for a `modelName` -*/}
      {/* (<IgntCrud storeName="OrgRepoModule" itemName="modelName" />) */}
      input: <input onChange={(event) => {
        const { name, parent } = getNameAndParent(event.target.value);
        setName(name);
        setParent(parent);
      }}></input>
      <p>{domain?.name}</p>
      <p>{domain?.parent}</p>
      <p>{domain?.owner}</p>
    </div>
  );
}
