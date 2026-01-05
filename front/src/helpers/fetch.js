function makeUrl(path) {
  return `http://localhost:8080${path}`
}

export async function fetchWithAuth(path, options = {}) {
  const res = await fetch(makeUrl(path), {
    ...options,
    credentials: 'include',
  })

  return res
}

export async function fetchWithRefresh(path, options = {}) {
  let res = await fetch(makeUrl(path), { ...options, credentials: 'include' })

  if (res.status === 401) {
    await fetch(makeUrl('/api/v1/auth/refresh'), { method: 'POST', credentials: 'include' })
    res = await fetch(makeUrl(path), { ...options, credentials: 'include' })
  }

  return res
}
