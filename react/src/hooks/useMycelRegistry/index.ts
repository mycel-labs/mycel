/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useMycelRegistry() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.MycelRegistry.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryDomain = (name: string, parent: string,  options: any) => {
    const key = { type: 'QueryDomain',  name,  parent };    
    return useQuery([key], () => {
      const { name,  parent } = key
      return  client.MycelRegistry.query.queryDomain(name, parent).then( res => res.data );
    }, options);
  }
  
  const QueryDomainAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryDomainAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.MycelRegistry.query.queryDomainAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryParams,QueryDomain,QueryDomainAll,
  }
}