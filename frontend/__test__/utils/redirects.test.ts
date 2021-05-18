import {
  redirectToDashboard,
  redirectToLandingPage,
  redirectToErrorPage,
} from '../../utils/redirects'

describe('test redirects', () => {
  it('test redirectToDashboard', () => {
    expect(redirectToDashboard()).toEqual({
      redirect: {
        permanent: false,
        destination: '/dashboard',
      },
    })
  })

  it('test redirectToLandingPage', () => {
    expect(redirectToLandingPage()).toEqual({
      redirect: {
        permanent: false,
        destination: '/',
      },
    })
  })

  it('test redirectToErrorPage', () => {
    const message = 'Network error'
    expect(redirectToErrorPage(message)).toEqual({
      redirect: {
        permanent: false,
        destination: `/error${message ? `?message=${message}` : ''}`,
      },
    })
  })
})
