import { Component, OnInit } from "@angular/core";
import { NzFormatEmitEvent } from "ng-zorro-antd/tree";
import { HttpClient, HttpErrorResponse } from "@angular/common/http";

declare var CodeMirror: any;

@Component({
  selector: "app-nz-demo-tree-draggable",
  templateUrl: "./config.component.html",
  styleUrls: ["./config.component.less"],
})
export class ConfigComponent implements OnInit {
  editor: any;
  constructor(private http: HttpClient) {}

  ngOnInit() {
    let url = "/api/core/v1";
    this.http.get(url).subscribe(
      (data: any) => {
        //  data:返回的数据
        console.log(data);
      },
      (error: HttpErrorResponse) => {
        console.log(error.status);
      }
    );
  }
  ngAfterViewInit() {
    this.editor = CodeMirror.fromTextArea(document.getElementById("code"), {
      lineNumbers: true,
      theme: "solarized",
      mode: "yaml",
      gutters: [
        "CodeMirror-linenumbers",
        "CodeMirror-foldgutter",
        "CodeMirror-lint-markers",
      ],
      lineWrapping: true,
      foldGutter: true,
      styleActiveLine: true,
    });
    this.editor.setSize("900px", "100%");
    this.editor.setValue = "123";
  }
  nodes = [
    {
      title: "0-0",
      key: "00",
      expanded: true,
      children: [
        {
          title: "0-0-0",
          key: "000",
          expanded: true,
          children: [
            { title: "0-0-0-0", key: "0000", isLeaf: true },
            { title: "0-0-0-1", key: "0001", isLeaf: true },
            { title: "0-0-0-2", key: "0002", isLeaf: true },
          ],
        },
        {
          title: "0-0-1",
          key: "001",
          children: [
            { title: "0-0-1-0", key: "0010", isLeaf: true },
            { title: "0-0-1-1", key: "0011", isLeaf: true },
            { title: "0-0-1-2", key: "0012", isLeaf: true },
          ],
        },
        {
          title: "0-0-2",
          key: "002",
        },
      ],
    },
    {
      title: "0-1",
      key: "01",
      children: [
        {
          title: "0-1-0",
          key: "010",
          children: [
            { title: "0-1-0-0", key: "0100", isLeaf: true },
            { title: "0-1-0-1", key: "0101", isLeaf: true },
            { title: "0-1-0-2", key: "0102", isLeaf: true },
          ],
        },
        {
          title: "0-1-1",
          key: "011",
          children: [
            { title: "0-1-1-0", key: "0110", isLeaf: true },
            { title: "0-1-1-1", key: "0111", isLeaf: true },
            { title: "0-1-1-2", key: "0112", isLeaf: true },
          ],
        },
      ],
    },
    {
      title: "0-2",
      key: "02",
      isLeaf: true,
    },
  ];
  nzEvent(event: NzFormatEmitEvent): void {
    console.log(event);
    console.log(event.keys);
    this.code = event.keys[0];
    this.editor.doc.setValue(event.keys[0]);
    console.log(this.editor);
  }
  code = "123";
}
