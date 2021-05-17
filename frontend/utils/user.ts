export const validateUser = (req: any): boolean => {
  return req.session.get('user')
  // TODO : add try catch
  // TODO: need to add validation
}
