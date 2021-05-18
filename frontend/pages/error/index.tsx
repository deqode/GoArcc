import Head from 'next/head'
import { useRouter } from 'next/router'
import { Fragment, ReactElement } from 'react'
import BasicLayout from '../../components/layouts/BasicLayout'

const ErrorPage = (): ReactElement => {
  const router = useRouter()
  return (
    <Fragment>
      <Head>
        <title>Error !!</title>
        <link rel="icon" href="assets/alfred.png" />
      </Head>
      <BasicLayout
        heading={'Alfred has encountered wit h'}
        component={router.query.message || 'Error'}
      />
    </Fragment>
  )
}

export default ErrorPage
