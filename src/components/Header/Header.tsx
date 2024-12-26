import '@/styles/globals.css';
import styles from './Header.module.css';
import { Heading, Spacer, Stack } from '@chakra-ui/react';

export default function Header() {
  return (
    <Stack className={styles.headerContainer}>
      <Spacer />
      <Stack direction={'row'}>
        <Spacer />
        <Stack direction={'row'} className={styles.headerContent}>
          <Heading size={'xl'}>DAkken Slide Maker</Heading>
        </Stack>
        <Spacer />
      </Stack>
      <Spacer />
    </Stack>
  );
}
