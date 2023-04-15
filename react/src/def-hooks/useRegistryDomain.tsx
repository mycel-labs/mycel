import { useState } from "react";
import { useClient } from "../hooks/useClient";
import { RegistryDomain } from "mycel-client-ts/mycel.registry/rest";
import { convertToNameAndParent } from "../utils/domainName";

export const useRegistryDomain = () => {
  const client = useClient();
  const [isLoading, setIsLoading] = useState(false);
  const [registryDomain, setRegistryDomain] = useState<RegistryDomain | null>(null)

  const updateRegistryDomain = async (domainName: string) => {
    const { name, parent } = convertToNameAndParent(domainName);
    setIsLoading(true);
    try {
      if (!name || !parent) {
        throw new Error("name or parent are empty")
      }
      const domain = await client.MycelRegistry.query.queryDomain(name, parent);
      setRegistryDomain(domain.data.domain || null);
    } catch (e) {
      console.error(e);
      setRegistryDomain(null);
      setIsLoading(false);
    }
    setIsLoading(false);
  }

  return {registryDomain, isLoading, updateRegistryDomain}
};