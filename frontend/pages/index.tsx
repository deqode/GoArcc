import { useQuery } from '@apollo/client';
import { Grid, Paper, Typography } from '@material-ui/core';
import Head from 'next/head';
import { GET_TOKEN } from '../GraphQL/cache';
import { SERVER } from '../utils/constants';
import Link from 'next/link'
import { report } from 'node:process';

export default function Login({ auth,repo }) {

  const { loading, error, data } = useQuery(GET_TOKEN);

  console.log(data, loading, error, "from index")

  return (
    <div>
      <Head>
        <title>Login to Alfred</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Paper elevation={0} style={{ padding: "100px" }} />
      <Grid justify="center" alignItems="center" spacing={1} container>
        <Grid item md={4}>
          <Typography variant="h1" component="h1" style={{ fontSize: "50px" }}>Let us make your cloud work for you</Typography>
          <Typography variant="subtitle2" style={{ fontSize: "15px" }}>provide your repo, select your cloud, and let Alfred do the heavy lifting.</Typography>

        </Grid>
        <Grid item>
          <div className="text-center">
            {data.token != null ?
              <> <div className="sign_up_head">Connect</div>
                <Link href={repo.redirectUrl}>
                  <a className="btn github_btn">Connect with github<img src="/assets/github_icon.png" alt="Login with github" /></a>
                </Link>  </> :
              <>
                <div className="sign_up_head">Sign Up</div>
                <a href={auth.url} className="btn github_btn">Login with github<img src="/assets/github_icon.png" alt="Login with github" /></a>
              </>}
          </div>                    </Grid>
      </Grid>

      <Paper />

    </div>
  )
}


export const getStaticProps = async () => {
  let res = await fetch(`${SERVER}/authentication/login`)
  let data = await res.json()
  let auth = data
  res = await fetch(`${SERVER}/vcs-connection/authorize/GITHUB`)
  data = await res.json()
  let repo = data
  return {
    props: {
      auth,
      repo
    },
  }
}