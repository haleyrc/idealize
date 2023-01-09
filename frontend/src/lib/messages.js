export async function getMessage() {
  const resp = await fetch("http://localhost:8080/api/GetMessage", {
    method: "POST",
  })
  const data = await resp.json()
  return data
}
