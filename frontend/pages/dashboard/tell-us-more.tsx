import { Paper } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import { withIronSession } from 'next-iron-session'
import Head from 'next/head'
import { ReactElement, useContext, useEffect } from 'react'
import { getOwnerName } from '../../api/gql/vcs'
import { getAllUserAccounts } from '../../api/rest/user'
import BasicLayout from '../../components/layouts/BasicLayout'
import CloneRepository from '../../components/tellUsMore/CloneRepository'
import ShowBranches from '../../components/tellUsMore/ShowBranches'
import ShowRepos from '../../components/tellUsMore/ShowRepos'
import useGetTellUsMoreState from '../../components/tellUsMore/useGetTellUsMoreState'
import UserContext from '../../contexts/UserContext'
import { UserResponse } from '../../interface'
import { sessionCongfig } from '../../utils/constants'
import { redirectToErrorPage, redirectToLandingPage } from '../../utils/redirects'
import { validateUser } from '../../utils/user'

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
  }, [user])

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
    <Paper elevation={0}>
      <Head>
        <title>Tell Us More!</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <BasicLayout
        heading={'Tell Us More'}
        subHeading={'You have so many cool projects.Let me know the one you want to deploy today '}
        // TODO:Transfer to new layout
        component={
          <Paper elevation={0}>
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
        }
      />
    </Paper>
  )
}
export const getServerSideProps = withSentry(
  withIronSession(async ({ req }) => {
    if (validateUser(req)) {
      const user = req.session.get('user')
      const res = await getAllUserAccounts(user.userId, user.idToken)

      if (res.error && res.accounts.length === 0) {
        return redirectToErrorPage('Network Error')
      }
      const resp = await getOwnerName({ idToken: user.idToken, accountId: res.accounts[0].id })
      if (resp.error) return redirectToLandingPage()
      return {
        props: {
          userID: user.userId,
          accountId: res.accounts[0].id,
          user,
          ownerName: resp.ownerName,
        },
      }
    }

    return redirectToLandingPage()
  }, sessionCongfig)
)

export default TellUsMore
