import { handler } from '../../../pages/api/session/destroy'

describe('destroy session tests', () => {
  it('should destroy session', async () => {
    const req = {
      session: { set: jest.fn(), destroy: jest.fn() },
      method: 'POST',
      body: { name: 'abc', email: 'abc@abc.com', password: 'abc123' },
    }
    const res = {
      status: jest.fn(() => res),
      end: jest.fn(),
      json: jest.fn(),
    }
    // TODO:need to handle cases
    const response = await handler(req, res)
    expect(response).toBeUndefined()
  })
})
