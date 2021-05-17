import { validateUser } from '../../utils/user'

describe('test User', () => {
  it('should return true', () => {
    const req = {
      session: { get: jest.fn(() => true) },
    }
    expect(validateUser(req)).toEqual(true)
  })
  it('should return false', () => {
    const req = {
      session: { get: jest.fn(() => false) },
    }
    expect(validateUser(req)).toEqual(false)
  })
})
