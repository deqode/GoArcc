import React, { useEffect } from 'react'
import { useRouter } from 'next/router';
import { SERVER, sessionCongfig } from '../../../utils/constants';
import { withIronSession } from 'next-iron-session';

function Callback(props:any) {
    const router = useRouter();
    const { user } = props

    const { query } = router
    useEffect(() => {
        if (query.code && query.code.length > 0) {
            (async () => {
                console.log(query)
                if (query.code && user.userId != "") {


                    let res = await fetch(`${SERVER}/account/get-user-all-account/${user.userId}`, {
                        headers: new Headers({
                            'Authorization': `Bearer ${user.idToken}`,
                        })
                    })
                    let data = await res.json()
                    const newUser = user

                    newUser.accounts = data.accounts || []

                    // todo: gql
                    res = await fetch(`${SERVER}/vcs-connection/GITHUB/callback?code=${query.code}&account_id=${newUser.accounts[0].id}`, {
                        headers: new Headers({
                            'Authorization': `Bearer ${user.idToken}`,
                        })
                    })
                    data = await res.json()
                    newUser.provider = data.provider

                    console.log(newUser)
                    const sessionRes = await fetch(`/api/session/set`, {
                        method: 'POST',
                        headers: {
                            'Accept': 'application/json',
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(newUser)
                    })
                    const sessionData = await sessionRes.json()
                    if (!sessionData) {
                        router.push("/dashboard")
                        return
                    }
                    router.push("/tell-us-more")
                } else {
                    router.push("/dashboard")
                }

            })()

        }
        // else
        // router.push("/")
    }, [query, router, user])

    return (
        <div>

        </div>
    )
}

export default Callback

export const getServerSideProps = withIronSession(async ({ req }) => {
    if (req.session.get("user")) {
        return { props: { user: req.session.get("user") } }
    }
    return {
        redirect: {
            permanent: false,
            destination: "/"
        }
    }

}, sessionCongfig)