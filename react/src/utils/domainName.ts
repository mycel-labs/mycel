
export const convertToNameAndParent = (domain: string) => {
  const s = domain.split(".");
  if (s.length === 1) {
    return { name: "", parent: s[0] };
  }
  return { name: s[0], parent: s.slice(1).join(".") }
}

export const convertToDomainName = (name: string | undefined, parent: string | undefined) => {
  return name && parent ? name + "." + parent : ""
}