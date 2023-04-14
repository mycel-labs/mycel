import { useMemo } from "react";
import useCosmosBankV1Beta1 from "../hooks/useCosmosBankV1Beta1";
import { useAddressContext } from "./addressContext";
import useMycelRegistry from "../hooks/useMycelRegistry";

export const useRegistryDomain = (name: string, parent: string) => {
  if (name === "" || parent === "") {
    return { domain: undefined, isLoading: true};
  }
  const { QueryDomain } = useMycelRegistry()
  const query = QueryDomain(name, parent, {});
  return { domain: query.data?.domain, isLoading: query.isLoading };
};