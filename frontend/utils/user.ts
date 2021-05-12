export const validateUser = (req: any): boolean => {
  if (!req.session.get('user')) {
    return false
  }
  return true
}
