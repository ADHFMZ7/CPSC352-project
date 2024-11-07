import { cookies } from 'next/headers'


export async function login(formdata: FormData) {

  const user = {username: formData.get('username')}

  // Create session
  // Expiration
  // Query database

  cookies().set('session', '', q

}
