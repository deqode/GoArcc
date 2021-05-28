import { Grid, Paper } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import { ReactElement } from 'react'

import { getOwnerName } from '../../api/gql/vcs'
import { getAllUserAccounts } from '../../api/rest/user'
import BasicLayout from '../../components/layouts/BasicLayout'
import PageHead, { Titles } from '../../components/PageHead'
import CloneRepository from '../../components/tellUsMore/CloneRepository'
import useGetTellUsMoreState from '../../components/tellUsMore/hooks/useGetTellUsMoreState'
import ShowBranches from '../../components/tellUsMore/ShowBranches'
import ShowRepos from '../../components/tellUsMore/ShowRepos'
import { IronSessionRequest, UserResponse } from '../../intefaces/interface'
import {
  redirectToDashboard,
  redirectToErrorPage,
  redirectToLandingPage,
} from '../../utils/redirects'
import { sessionPropsWrapper, validateUser } from '../../utils/user'

interface TellUsMoreProps {
  user: UserResponse
  accountId: string
  ownerName: string
}
const TellUsMore = ({ user, accountId, ownerName }: TellUsMoreProps): ReactElement => {
  const provider = 'GITHUB'
  const { state, dispatch, setCloneData } = useGetTellUsMoreState()
  const { branchName, cloneUrl, currentRepo } = state
  return (
    <Paper elevation={0}>
      <PageHead title={Titles.DASHBOARD} />
      <BasicLayout
        heading={'Tell Us More'}
        subHeading={'You have so many cool projects.Let me know the one you want to deploy today '}
        // TODO:Transfer to new layout
        component={
          <Grid container alignItems="flex-start" spacing={2}>
            <ShowRepos
              userId={user.userId}
              accountId={accountId}
              setCurrentRepo={dispatch}
              provider={provider}
            />
            <ShowBranches
              ownerName={ownerName}
              repoName={currentRepo}
              accountId={accountId}
              provider={provider}
              set={dispatch}
            />
            <CloneRepository
              ownerName={ownerName}
              accountId={accountId}
              provider={provider}
              branchName={branchName}
              cloneURL={cloneUrl}
              setCloneData={setCloneData}
            />
          </Grid>
        }
      />
    </Paper>
  )
}
export const handler = async ({ req }: { req: IronSessionRequest }) => {
  if (validateUser(req)) {
    const user = req.session.get('user')
    const responseGetAllUserAccounts = await getAllUserAccounts(user.userId, user.idToken)
    if (responseGetAllUserAccounts.error && !responseGetAllUserAccounts.accounts.length) {
      return redirectToErrorPage('Network Error')
    }
    const responseGetOwnerName = await getOwnerName({
      idToken: user.idToken,
      accountId: responseGetAllUserAccounts.accounts[0].id,
    })

    if (responseGetOwnerName.error) return redirectToDashboard()
    return {
      props: {
        userID: user.userId,
        accountId: responseGetAllUserAccounts.accounts[0].id,
        user,
        ownerName: responseGetOwnerName.ownerName,
      },
    }
  }
  return redirectToLandingPage()
}
export const getServerSideProps = withSentry(sessionPropsWrapper(handler))

export default TellUsMore
