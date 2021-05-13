import { Grid, makeStyles } from '@material-ui/core'
import React, { ReactElement } from 'react'
const useStyles = makeStyles({
  root: {
    paddingTop: 100,
  },
})
const BasicLayout = ({
  heading,
  component,
}: {
  heading: ReactElement | string
  component: ReactElement | string
}): ReactElement => {
  const classes = useStyles()
  return (
    <Grid container className={classes.root}>
      <Grid item xs={6} sm={2} md={8} lg={12} xl={12}>
        {heading}
      </Grid>
      <Grid item xs={6} sm={10} md={4} lg={12} xl={12}>
        {component}
      </Grid>
    </Grid>
  )
}

export default BasicLayout
