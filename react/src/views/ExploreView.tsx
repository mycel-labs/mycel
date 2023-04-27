import { useRef, useState, useEffect } from "react";
import { useClient } from "../hooks/useClient";
import { RegistryDomain } from "mycel-domain-mycel-client-ts/mycel.registry/rest";
import { IgntButton } from "@ignt/react-library";

import Fuse from "fuse.js";

interface FuseDomainEntry {
  label: string;
  value: RegistryDomain;
}

export default function ExploreView() {
  const client = useClient();
  const [query, setQuery] = useState<string>("");
  const [domains, setDomains] = useState<FuseDomainEntry[]>([]);
  const [result, setResult] = useState<RegistryDomain[]>([]);
  const fuse = useRef<Fuse<FuseDomainEntry> | null>(null);

  const getDomainList = async () => {
    const response = await client.MycelRegistry.query.queryDomainAll();
    setDomains(
      (response.data.domain ?? []).map((e) => ({
        label: e.name + "." + e.owner,
        value: e,
      })),
    );
  };

  useEffect(() => {
    getDomainList();
  }, []);

  useEffect(() => {
    fuse.current = new Fuse(domains, {
      keys: ["label"],
      includeScore: true,
    });
    setResult(domains.map((e) => e.value));
  }, [domains]);

  useEffect(() => {
    if (!query) {
      setResult(domains.map((e) => e.value));
      return;
    }
    const results = fuse.current?.search(query);
    setResult(results?.map(({ item }) => item.value) ?? []);
    console.log(result);
  }, [query]);

  useEffect(() => {
    console.log(result);
  }, [result]);

  return (
    <div className="w-3/4 mx-auto">
      <h2 className=" text-2xl">Explore Domain</h2>
      <div className="flex mt-2 p-2 justify-between">
        <input
          className="mr-6 mt-1 py-2 px-4 h-14 bg-gray-100 w-full border-xs text-base leading-tight rounded-xl outline-0"
          placeholder="Mycel Domain"
          onChange={(event) => {
            setQuery(event.target.value);
          }}
        />
      </div>
      <div>
        {result.map((e) => (
          <div className="w-full flex justify-between my-4" key={e.name + "." + e.parent}>
            <h2 className=" text-2xl m-2 font-semibold">{e.name + "." + e.parent}</h2>
            <a href={"/resolve?name=" + e.name + "&parent=" + e.parent}>
              <IgntButton className="mt-1 h-10 w-48">Resolve</IgntButton>
            </a>
          </div>
        ))}
      </div>
    </div>
  );
}
