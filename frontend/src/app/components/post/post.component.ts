import { Component, Input } from '@angular/core';

interface postData {
  id: string,
  title: string,
  content: string,
};

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.scss']
})
export class PostComponent {
  @Input() postData: postData;
  
  constructor() {

  }
}
