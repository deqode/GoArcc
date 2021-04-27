import { useQuery } from '@apollo/client';
import { AppBar, Button, Grid, makeStyles, Toolbar, Typography } from '@material-ui/core'
import Link from 'next/link'

import React from 'react'
import { GET_TOKEN } from '../GraphQL/cache';

function Navbar() {
    const { loading, error, data } = useQuery(GET_TOKEN);

    return (
        <AppBar position="static">
            <Toolbar>
                <Grid justify="space-between" container>
                    <Grid item>
                        <Link href="/">
                            <img src="/assets/alfred.png" width="100px" />
                        </Link>
                    </Grid>
                    <Grid item>
                        {data.token != null ?
                            <Button color="inherit">Logout</Button>
                            : <Button color="inherit">Login</Button>}
                    </Grid>
                </Grid>
            </Toolbar>
        </AppBar >
    )
}

export default Navbar
