import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Input,
} from "@chakra-ui/react";
import { FormEventHandler } from "react";

function Register() {
  const onSubmit = (e: any) => {
    e.preventDefault();

    const { email, password, password_confirmation, username } =
      e.target.elements;

    fetch("/api/auth/register", {
      body: JSON.stringify({
        email: email.value,
        password: password.value,
        password_confirmation: password_confirmation.value,
        username: username.value,
      }),
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      method: "POST",
      credentials: "include",
    })
      .then((res) => res.json())
      .then(console.log)
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <Box mt={8}>
      <Container>
        <form onSubmit={onSubmit}>
          <FormControl isRequired>
            <FormLabel htmlFor="username">Username</FormLabel>
            <Input
              id="username"
              name="username"
              placeholder="Username"
              required
            />
          </FormControl>

          <FormControl isRequired mt={4}>
            <FormLabel htmlFor="email">Email</FormLabel>
            <Input
              id="email"
              name="email"
              placeholder="Email"
              type="email"
              required
            />
          </FormControl>

          <FormControl isRequired mt={4}>
            <FormLabel htmlFor="password">Password</FormLabel>
            <Input
              id="password"
              name="password"
              placeholder="Password"
              type="password"
              required
            />
          </FormControl>

          <FormControl isRequired mt={4}>
            <FormLabel htmlFor="password_confirmation">
              Password Confirmation
            </FormLabel>
            <Input
              id="password_confirmation"
              name="password_confirmation"
              placeholder="Password confirmation"
              type="password"
              required
            />
          </FormControl>

          <Button mt={8} colorScheme="blue" w="100%" type="submit">
            Submit
          </Button>
        </form>
      </Container>
    </Box>
  );
}

export default Register;