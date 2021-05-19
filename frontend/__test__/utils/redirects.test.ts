import {
  redirectToDashboard,
  redirectToLandingPage,
  redirectToErrorPage,
} from '../../utils/redirects'

describe('redirect utils test', () => {
  it('should redirect to dashboard', () => {
    expect(redirectToDashboard()).toEqual({
      redirect: {
        permanent: false,
        destination: '/dashboard',
      },
    })
  })

  it('should redirect to landing page', () => {
    expect(redirectToLandingPage()).toEqual({
      redirect: {
        permanent: false,
        destination: '/',
      },
    })
  })

  it('should redirect to error page', () => {
    const message = 'Network error'
    expect(redirectToErrorPage(message)).toEqual({
      redirect: {
        permanent: false,
        destination: `/error${message ? `?message=${message}` : ''}`,
      },
    })
  })
})
