import { Stack } from '@chakra-ui/react';
import { Providers } from './providers';
import '@/styles/globals.css';
import Header from '@/components/Header/Header';
import Main from '@/components/Main/Main';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html>
      <body>
        <Providers>
          <Stack>
            <Header />
            <Main>{children}</Main>
          </Stack>
        </Providers>
      </body>
    </html>
  );
}
