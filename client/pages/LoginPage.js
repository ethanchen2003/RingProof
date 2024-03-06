import React, { useState } from 'react';
import { View, TextInput, Button, Text, StyleSheet, Pressable } from 'react-native';

const LoginPage = ({ onLoginSuccess }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

  const handleSubmit = async () => {
    try {
      const response = await fetch('http://100.64.19.138:8080/signin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });
      if (response.ok) {
        onLoginSuccess();
      } else {
      }
    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  const handleForgotPasswordPress = () => {
    // Handle the press on Forgot Password link
  };

  const handleSignInPress = () => {
    // Handle the press on Sign In link
  };

  const handleContactUsPress = () => {
    // Handle the press on Contact Us link
  };

  return (
    <View style={styles.container}>
      <View style={styles.loginForm}>
        <Text style={styles.companyName}>RingProof</Text>
        <TextInput
          style={styles.input}
          placeholder="Username"
          placeholderTextColor={styles.inputPlaceholder.color} 
          required
          value={username}
          onChangeText={setUsername}
        />
        <TextInput
          style={styles.input}
          placeholder="Password"
          placeholderTextColor={styles.inputPlaceholder.color} 
          required
          value={password}
          onChangeText={setPassword}
          secureTextEntry
        />
        <Pressable style={styles.loginButton} onPress={handleSubmit}>
          <Text style={styles.buttonText}>LOGIN</Text>
        </Pressable>
      </View>
      <View style={styles.footerLinksContainer}>
      <View style={styles.footerLinksRow}>
        <Pressable onPress={handleForgotPasswordPress}>
          <Text style={styles.linkText}>Forgot Password?</Text>
        </Pressable>
        <Pressable onPress={handleSignInPress}>
          <Text style={styles.linkText}>Sign In</Text>
        </Pressable>
      </View>
      <Pressable onPress={handleContactUsPress}>
        <Text style={styles.linkTextCentered}>Contact us!</Text>
      </Pressable>
    </View>
    </View>
  );
};
const styles = StyleSheet.create({
    container: {
      flex: 1,
      justifyContent: 'center',
      alignItems: 'center',
      backgroundColor: '#0D1B2A',
    },
    loginForm: {
      width: 300,
      padding: 40,
      backgroundColor: '#1B263B',
      borderRadius: 8,
      alignItems: 'center',
      elevation: 10,
      shadowColor: '#000000',
      shadowOffset: { width: 0, height: 4 },
      shadowOpacity: 0.5, 
      shadowRadius: 10,
    },
    companyName: {
      fontSize: 24,
      color: '#FFFFFF',
      marginBottom: 20,
      fontWeight: 'bold',
    },
    input: {
      width: '100%',
      height: 26,
      paddingHorizontal: 10,
      marginVertical: 10,
      borderRadius: 20,
      backgroundColor: '#14213D',
      color: '#FFFFFF',
      fontSize: 18,
    },
    inputPlaceholder: {
      color: '#A2A2A2',
    },
    loginButton: {
        marginTop: 20,
        backgroundColor: '#0A79DF',
        borderRadius: 20,
        paddingVertical: 10,
        paddingHorizontal: 20,
        width: '100%',
        alignItems: 'center',
        justifyContent: 'center',
    },
    buttonText: {
      color: '#FFFFFF', 
      fontSize: 18,
    },
    footerLinksContainer: {
        width: 500,
        alignItems: 'center',
        marginTop: 20,
    },
    footerLinksRow: {
        flexDirection: 'row',
        justifyContent: 'space-evenly',
        width: '100%',
    },
    linkText: {
        color: 'lightblue',
        marginVertical: 8,
    },
    linkTextCentered: {
        color: 'lightblue',
        marginTop: 16,
        textAlign: 'center',
    },
  });
export default LoginPage;
