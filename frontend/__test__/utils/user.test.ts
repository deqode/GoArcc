import { validateUser } from '../../utils/user'

describe('test if user session available', () => {
  it('should return true })', () => {
    const req = {
      session: { get: jest.fn(() => true) },
    }
    expect(validateUser(req)).toEqual(true)
  })
  it('should return false if user session not available', () => {
    const req = {
      session: { get: jest.fn(() => false) },
    }
    expect(validateUser(req)).toEqual(false)
  })
})
