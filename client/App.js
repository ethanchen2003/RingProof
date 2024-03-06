import React, { useState } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import LoginPage from './pages/LoginPage';
import WelcomePage from './pages/WelcomePage';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  return (
    <NavigationContainer>
      {isLoggedIn ? <WelcomePage /> : <LoginPage onLoginSuccess={() => setIsLoggedIn(true)} />}
    </NavigationContainer>
  );
}

export default App;


