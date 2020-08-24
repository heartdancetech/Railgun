import { NgModule } from '@angular/core';

import { ConfigRoutingModule } from './config-routing.module';

import { ConfigComponent } from './config.component';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzTreeModule } from 'ng-zorro-antd/tree';

@NgModule({
  imports: [ConfigRoutingModule, NzButtonModule, NzTreeModule],
  declarations: [ConfigComponent],
  exports: [ConfigComponent],
})
export class ConfigModule {}
