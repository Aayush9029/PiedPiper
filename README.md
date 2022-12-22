# PiedPiper

_Middle to Out_

### 😠 Problem

- I have a home server and several Internet-connected nodemcu devices at different locations.
- I want to be able to access data collected by these devices from anywhere in the world using my Phone or Watch.

![CleanShot 2022-12-21 at 20 38 37](https://user-images.githubusercontent.com/43297314/209244569-533e0080-0cc3-4464-b777-3badf2952280.png)

### 🤓 Techincal Problem

0. _Access the internal data.._
1. Without having to open a port in my home router.
2. I can't open ports on other locations.
3. Without having to use ssh tunneling or reverse proxy on home router.
4. I can't use ssh tunneling or reverse proxy on other devices.
5. Without having to setup a static IP.
6. Without turning my client device into a server and somehow connecting my Devices to client directly (also, see 2, 4)

### 💡 Solution

- Send a POST request with key to a remote server to add / update data.

```sh
$ curl -X POST yourserver.url/super-secret-key -d 'I DID DONE IT... 20MB worth JSON b64 encoded data...'
```

- Send a GET request with the same key to get the data.

```sh
$ curl -X GET yourserver.url/super-secret-key
$ I DID DONE IT... 20MB worth JSON b64 encoded data...
```

---

### ✨ Features

- Data persistence
- Rate Limiting
- Monitor API + Server status
- REST API for adding, updating and getting the data


✅ Deployed and tested 

<img src="https://user-images.githubusercontent.com/43297314/209243866-d65139a9-43f3-499f-92ec-4e25269e0463.png" width="512">
