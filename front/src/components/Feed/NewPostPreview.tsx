import { JSXElement } from 'solid-js';
import { Button } from '~/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';
import { Separator } from '~/components/ui/separator';

interface NewPostPreviewProps {
  body?: string;
}

export default function NewPostPreview(props: NewPostPreviewProps): JSXElement {
  return (
    <Dialog defaultOpen={true}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Apparently he is Disappointed</DialogTitle>
          <DialogDescription>
            It's okay <b>clown</b> 🤡 don't cry yet. Wait till you read the
            comments
          </DialogDescription>
        </DialogHeader>
        <Separator />
        <DialogFooter>
          <Button variant={'secondary'}>Cancel</Button>
          <Button>Post</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
