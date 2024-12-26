import '@/styles/globals.css';
import { Stack } from '@chakra-ui/react';

export default function Main({ children }: { children: React.ReactNode }) {
  return <Stack>{children}</Stack>;
}
