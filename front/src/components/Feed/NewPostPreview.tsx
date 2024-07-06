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
import { AspectRatio } from '../ui/aspect-ratio';

interface NewPostPreviewProps {
  open: boolean;
  setOpen: (open: boolean) => void;
}

export default function NewPostPreview(props: NewPostPreviewProps): JSXElement {
  return (
    <Dialog open={props.open} onOpenChange={props.setOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Apparently he is Disappointed</DialogTitle>
          <DialogDescription>
            It's okay <b>clown</b> 🤡 don't cry yet. Wait till you read the
            comments
          </DialogDescription>
        </DialogHeader>

        <AspectRatio ratio={1 / 1} class='rounded bg-muted'>
          <Button variant={'secondary'} class='h-full w-full flex-col'>
            Upload an image
            <p class='font-light text-muted-foreground'>
              we recommend a resloution of 500x500 px
            </p>
          </Button>
        </AspectRatio>

        {/* <Button variant={'secondary'} class=''>
          Pick an image
        </Button> */}

        <Separator />
        <DialogFooter>
          <Button variant={'secondary'}>Cancel</Button>
          <Button>Post</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
