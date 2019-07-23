declare module 'md5' {
  function md5(data: string, options?: {encoding: string, asBytes?: boolean, asString?: boolean}): string
  namespace md5 {}
  export = md5
}
