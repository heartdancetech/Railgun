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

  ngOnInit() {}

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
  }
  key = "";
  title = "title";
  isLoadingSave = false;
  nodes = [
    {
      title: "/conf/",
      key: "/conf/",
      selected: false,
      selectable: false,
      disabled: true,
      expanded: false,
    },
  ];
  nzEvent(event: NzFormatEmitEvent): void {
    if (event.eventName === "expand") {
      if (event.node.isExpanded) {
        event.node.clearChildren();
        this.key = "";
        this.title = "title";
      }
      this.getNode().then((node) => {
        event.node.addChildren(node);
      });
    }
    if (event.eventName === "click") {
      this.getNodeValue(event.keys[0])
        .then((data) => {
          this.title = event.keys[0];
          this.key = event.keys[0];
          this.editor.doc.setValue(data);
        })
        .catch((err) => {
          console.log(err);
        });
    }
  }

  getNode(): Promise<any> {
    let url = "/api/conf/get_keys";
    return new Promise((resolve, reject) => {
      this.http.get(url).subscribe(
        (data: any) => {
          let node = [];
          data["data"].forEach((v: string, i: number) => {
            let vs = v.split("/conf/");
            node[i] = { title: vs[1], key: v, isLeaf: true };
          });

          resolve(node);
        },
        (error: HttpErrorResponse) => {
          reject(error);
        }
      );
    });
  }

  getNodeValue(key: string): Promise<string> {
    let url = `/api/conf/get_value?key=${key}`;
    return new Promise((resolve, reject) => {
      this.http.get(url).subscribe(
        (data: any) => {
          resolve(data["data"]);
        },
        (error: HttpErrorResponse) => {
          reject(error);
        }
      );
    });
  }

  onClickSave() {
    console.log(this.editor.doc.getValue());
    console.log(this.key);
    this.isLoadingSave = true;
    let url = "/api/conf/save_value";
    const body = {
      key: this.key,
      value: this.editor.doc.getValue(),
    };
    this.http.put(url, body).subscribe(
      () => {},
      (err) => {
        console.log(err);
      },
      () => {
        this.isLoadingSave = false;
      }
    );
  }
}
