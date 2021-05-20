import { Grid, makeStyles, Paper, Typography } from '@material-ui/core'
import React, { ReactElement } from 'react'

const useStyles = makeStyles((theme) => ({
  root: {
    paddingTop: 100,
    paddingBottom: 30,
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
  },
  text: {
    display: 'inline-block',
  },
}))
const BasicLayout = ({
  heading,
  subHeading,
  component,
}: {
  heading: string
  subHeading?: string
  component: ReactElement | string | string[]
}): ReactElement => {
  const classes = useStyles()
  return (
    <Grid
      container
      direction="row"
      justify="space-around"
      alignItems="center"
      className={classes.root}>
      <Grid item xs={12} sm={6} md={4} lg={4} xl={4}>
        <Paper className={classes.paper} elevation={0}>
          <Typography color="primary" variant="h3" component="h1">
            {heading}
          </Typography>
          <br />
          <Typography color="secondary" variant="h5">
            {subHeading}
          </Typography>
        </Paper>
      </Grid>
      <Grid item xs={12} sm={6} md={4} lg={4} xl={4}>
        <Paper className={classes.paper} elevation={0}>
          <Typography display="inline" className={classes.text}>
            {component}
          </Typography>
        </Paper>
      </Grid>
    </Grid>
  )
}

export default BasicLayout
