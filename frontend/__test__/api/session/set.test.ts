import { handler } from '../../../pages/api/session/set'

describe('set session tests', () => {
  it('should set session', async () => {
    const req = {
      session: { set: jest.fn(), save: jest.fn() },
      method: 'POST',
      body: { name: 'abc', email: 'abc@abc.com', password: 'abc123' },
    }
    const res = {
      status: jest.fn(() => res),
      end: jest.fn(),
      json: jest.fn(),
    }
    // TODO:need to handel cases
    const response = await handler(req, res)
    expect(response).toBeUndefined()
  })
})
