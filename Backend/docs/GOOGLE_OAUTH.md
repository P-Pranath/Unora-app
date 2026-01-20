# Google OAuth Flow for Mobile Apps

This document explains how to integrate Google OAuth in your mobile app with this backend.

## Overview

The flow uses **Google Sign-In ID tokens** from the mobile app:

```
Mobile App                    Backend                    Google
    │                            │                          │
    │──── Google Sign-In ────────┼─────────────────────────►│
    │◄─── ID Token ──────────────┼──────────────────────────│
    │                            │                          │
    │──── POST /auth/google ────►│                          │
    │     {idToken: "..."}       │                          │
    │                            │──── Verify Token ───────►│
    │                            │◄─── User Info ───────────│
    │                            │                          │
    │◄─── {accessToken,          │                          │
    │      refreshToken,         │                          │
    │      user}                 │                          │
    │                            │                          │
```

---

## Mobile App Setup

### iOS (Swift)

```swift
import GoogleSignIn

// Configure in AppDelegate
GIDSignIn.sharedInstance.configuration = GIDConfiguration(
    clientID: "YOUR_IOS_CLIENT_ID.apps.googleusercontent.com"
)

// Sign in
func signInWithGoogle() {
    GIDSignIn.sharedInstance.signIn(withPresenting: self) { result, error in
        guard error == nil, let user = result?.user else { return }
        
        // Get ID token to send to backend
        let idToken = user.idToken?.tokenString
        
        // Send to backend
        authenticateWithBackend(idToken: idToken!)
    }
}

func authenticateWithBackend(idToken: String) {
    let url = URL(string: "https://your-api.com/api/v1/auth/google")!
    var request = URLRequest(url: url)
    request.httpMethod = "POST"
    request.setValue("application/json", forHTTPHeaderField: "Content-Type")
    request.httpBody = try? JSONEncoder().encode(["idToken": idToken])
    
    URLSession.shared.dataTask(with: request) { data, _, error in
        // Handle response: save tokens, navigate to home
    }.resume()
}
```

### Android (Kotlin)

```kotlin
// Configure in build.gradle
implementation 'com.google.android.gms:play-services-auth:20.7.0'

// Sign in
private fun signInWithGoogle() {
    val gso = GoogleSignInOptions.Builder(GoogleSignInOptions.DEFAULT_SIGN_IN)
        .requestIdToken("YOUR_WEB_CLIENT_ID.apps.googleusercontent.com")
        .requestEmail()
        .build()
    
    val googleSignInClient = GoogleSignIn.getClient(this, gso)
    startActivityForResult(googleSignInClient.signInIntent, RC_SIGN_IN)
}

override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
    if (requestCode == RC_SIGN_IN) {
        val task = GoogleSignIn.getSignedInAccountFromIntent(data)
        val account = task.getResult(ApiException::class.java)
        
        // Get ID token to send to backend
        val idToken = account.idToken
        authenticateWithBackend(idToken!!)
    }
}

private fun authenticateWithBackend(idToken: String) {
    // Use Retrofit or similar
    val request = GoogleAuthRequest(idToken = idToken)
    api.googleLogin(request).enqueue(...)
}
```

### React Native

```javascript
import { GoogleSignin } from '@react-native-google-signin/google-signin';

// Configure
GoogleSignin.configure({
  webClientId: 'YOUR_WEB_CLIENT_ID.apps.googleusercontent.com',
  offlineAccess: true,
});

// Sign in
async function signInWithGoogle() {
  try {
    await GoogleSignin.hasPlayServices();
    const userInfo = await GoogleSignin.signIn();
    
    // Get ID token
    const { idToken } = await GoogleSignin.getTokens();
    
    // Send to backend
    const response = await fetch('https://your-api.com/api/v1/auth/google', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ idToken }),
    });
    
    const data = await response.json();
    // Save tokens: data.accessToken, data.refreshToken
  } catch (error) {
    console.error(error);
  }
}
```

### Flutter

```dart
import 'package:google_sign_in/google_sign_in.dart';

final GoogleSignIn _googleSignIn = GoogleSignIn(
  scopes: ['email', 'profile'],
);

Future<void> signInWithGoogle() async {
  final GoogleSignInAccount? account = await _googleSignIn.signIn();
  final GoogleSignInAuthentication auth = await account!.authentication;
  
  // Get ID token
  final String? idToken = auth.idToken;
  
  // Send to backend
  final response = await http.post(
    Uri.parse('https://your-api.com/api/v1/auth/google'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({'idToken': idToken}),
  );
  
  final data = jsonDecode(response.body);
  // Save: data['accessToken'], data['refreshToken']
}
```

---

## API Endpoints

### Login: `POST /api/v1/auth/google`

**Request:**
```json
{
  "idToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "accessToken": "encrypted-access-token",
    "refreshToken": "encrypted-refresh-token",
    "expiresIn": 86400,
    "tokenType": "Bearer",
    "user": {
      "id": "uuid",
      "email": "user@gmail.com",
      "name": "John Doe",
      "firstName": "John",
      "lastName": "Doe",
      "picture": "https://...",
      "provider": "google",
      "createdAt": "2024-01-01T00:00:00Z"
    }
  }
}
```

### Refresh Token: `POST /api/v1/auth/refresh`

**Request:**
```json
{
  "refreshToken": "encrypted-refresh-token"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "accessToken": "new-encrypted-access-token",
    "expiresIn": 86400,
    "tokenType": "Bearer"
  }
}
```

### Logout: `POST /api/v1/auth/logout`

**Headers:**
```
Authorization: Bearer <accessToken>
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "message": "Logged out successfully"
  }
}
```

### Get Current User: `GET /api/v1/auth/me`

**Headers:**
```
Authorization: Bearer <accessToken>
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "userId": "uuid",
    "email": "user@gmail.com",
    "provider": "google"
  }
}
```

---

## Token Management in Mobile App

### Store Tokens Securely

**iOS:** Use Keychain
```swift
KeychainWrapper.standard.set(accessToken, forKey: "accessToken")
KeychainWrapper.standard.set(refreshToken, forKey: "refreshToken")
```

**Android:** Use EncryptedSharedPreferences
```kotlin
val sharedPreferences = EncryptedSharedPreferences.create(...)
sharedPreferences.edit().putString("accessToken", token).apply()
```

### Using Access Token

Add to every authenticated request:
```
Authorization: Bearer <accessToken>
```

### Refresh Flow

```javascript
async function fetchWithAuth(url, options = {}) {
  const accessToken = await getAccessToken();
  
  const response = await fetch(url, {
    ...options,
    headers: {
      ...options.headers,
      'Authorization': `Bearer ${accessToken}`,
    },
  });
  
  // If 401, try refresh
  if (response.status === 401) {
    const newToken = await refreshAccessToken();
    if (newToken) {
      return fetch(url, {
        ...options,
        headers: {
          ...options.headers,
          'Authorization': `Bearer ${newToken}`,
        },
      });
    }
    // Refresh failed, logout user
    await logout();
  }
  
  return response;
}

async function refreshAccessToken() {
  const refreshToken = await getRefreshToken();
  
  const response = await fetch('/api/v1/auth/refresh', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ refreshToken }),
  });
  
  if (response.ok) {
    const data = await response.json();
    await saveAccessToken(data.data.accessToken);
    return data.data.accessToken;
  }
  
  return null;
}
```

---

## Environment Configuration

Add to `.env`:
```env
# JWT settings (adjust as needed)
JWT_ACCESS_TOKEN_EXPIRY_HOURS=24
JWT_REFRESH_TOKEN_EXPIRY_DAYS=30

# Auth secret (generate with: openssl rand -hex 32)
AUTH_SECRET_KEY=your-secure-random-key-here
```

---

## Security Notes

1. **ID Token Validation**: Backend verifies the Google ID token with Google's servers
2. **Token Encryption**: Access/refresh tokens are encrypted with ChaCha20-Poly1305
3. **Refresh Token Storage**: Stored in Redis for revocation capability
4. **Logout**: Revokes refresh token, invalidating all sessions

---

## Troubleshooting

| Error | Cause | Solution |
|-------|-------|----------|
| "Invalid token issuer" | Token not from Google | Verify Google Sign-In setup |
| "Token expired" | Access token expired | Use refresh token |
| "Refresh token revoked" | User logged out | Re-authenticate |
| "Invalid token" | Malformed or tampered | Re-authenticate |
