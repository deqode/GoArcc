import Head from 'next/head'
import { useRouter, withRouter } from 'next/router'
import { useEffect, useState } from 'react'

const ErrorPage = () => {
  const router = useRouter()
  const [message, setMessage] = useState('')
  useEffect(() => {
    if (router.query.message && typeof router.query.message === 'string')
      setMessage(router.query.message)
  }, [router])
  return (
    <div>
      <Head>
        <title>Error !!</title>
        <link rel="icon" href="assets/alfred.png" />
      </Head>
      <section className="content_section">
        <div className="container">
          <div className="row">
            <div className="col-md-6">
              <h1 className="page-head">Alfred has encountered with error</h1>
            </div>
            <div className="col-md-6">
              <h1 className="page-head">{message || ''}</h1>
            </div>
            <div className="col-md-6 my-auto login-right"></div>
          </div>
        </div>
      </section>
    </div>
  )
}

export default withRouter(ErrorPage)
