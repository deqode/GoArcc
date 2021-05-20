import Head from 'next/head'
import React from 'react'

export enum Titles {
  WELCOME = 'Welcome to Alfred',
  DASHBOARD = 'Dashboard',
  LOADING = 'Loading...',
  ERROR = 'Error',
  CUSTOM = 'custom',
}

const PageHead = ({ title }: { title: Titles }): JSX.Element => {
  return (
    <Head>
      <meta name="viewport" content="minimum-scale=1, initial-scale=1, width=device-width" />
      <title>{title}</title>
    </Head>
  )
}
export default PageHead
