import { useRouter } from 'next/router'
import { Fragment, ReactElement } from 'react'

import BasicLayout from '../../components/layouts/BasicLayout'
import PageHead, { Titles } from '../../components/PageHead'

const ErrorPage = (): ReactElement => {
  const router = useRouter()
  return (
    <Fragment>
      <PageHead title={Titles.ERROR} />
      <BasicLayout
        heading={'Alfred has encountered with'}
        component={router.query.message || 'Error'}
      />
    </Fragment>
  )
}

export default ErrorPage
