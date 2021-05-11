import { Paper } from '@material-ui/core'
import { withIronSession } from 'next-iron-session'
import Head from 'next/head'
import { ReactElement, useContext, useEffect } from 'react'
import { getOwnerName } from '../api/gql/vcs'
import { getAllUserAccounts } from '../api/rest/user'
import CloneRepository from '../components/tellUsMore/CloneRepository'
import ShowBranches from '../components/tellUsMore/ShowBranches'
import ShowRepos from '../components/tellUsMore/ShowRepos'
import useGetTellUsMoreState from '../components/tellUsMore/useGetTellUsMoreState'
import UserContext from '../contexts/UserContext'
import { UserResponse } from '../interface'
import { sessionCongfig } from '../utils/constants'
import { validateUser } from '../utils/user'

const TellUsMore = ({
  user,
  accountId,
  ownerName,
}: {
  user: UserResponse
  accountId: string
  ownerName: string
}): ReactElement => {
  const { setUser } = useContext(UserContext)
  const provider = 'GITHUB'
  useEffect(() => {
    if (user.idToken !== '') {
      setUser({
        loggedIn: true,
        idToken: user.idToken,
      })
    }
  }, [])

  const {
    currentRepo,
    setCurrentRepo,
    branchName,
    setBranchName,
    cloneUrl,
    setCloneUrl,
    setCloneData,
  } = useGetTellUsMoreState()

  return (
    <Paper>
      <Head>
        <title>Tell Us More!</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <ShowRepos
        userId={user.userId}
        accountId={accountId}
        setCurrentRepo={setCurrentRepo}
        provider={provider}
      />
      <ShowBranches
        ownerName={ownerName}
        repoName={currentRepo}
        accountId={accountId}
        provider={provider}
        setBranchName={setBranchName}
        setCloneUrl={setCloneUrl}
      />
      <CloneRepository
        ownerName={ownerName}
        accountId={accountId}
        provider={provider}
        branchName={branchName}
        cloneURL={cloneUrl}
        setCloneData={setCloneData}
      />
    </Paper>
  )
}
export const getServerSideProps = withIronSession(async ({ req }) => {
  if (validateUser(req)) {
    const user = req.session.get('user')
    const res = await getAllUserAccounts(user.userId, user.idToken)

    if (res.error && res.accounts.length === 0) {
      return {
        redirect: {
          permanent: false,
          destination: '/error',
        },
      }
    }
    const resp = await getOwnerName({ idToken: user.idToken, accountId: res.accounts[0].id })
    if (resp.error)
      return {
        redirect: {
          permanent: false,
          destination: '/',
        },
      }
    return {
      props: {
        userID: user.userId,
        accountId: res.accounts[0].id,
        user,
        ownerName: resp.ownerName,
      },
    }
  }

  return {
    redirect: {
      permanent: false,
      destination: '/',
    },
  }
}, sessionCongfig)

export default TellUsMore
