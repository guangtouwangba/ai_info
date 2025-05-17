# #RISC-V虚拟化实战##基于RISC-V手撸Hypervisor#如何基于RISC-V H-extension，手撸一个Hypervisor，并成功运行Linux？这位开发者在自研操作系统Starina实现虚拟化...

**URL**: https://weibo.com/6105753431/PsibZnknt

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23RISC-V%E8%99%9A%E6%8B%9F%E5%8C%96%E5%AE%9E%E6%88%98%23&amp;extparam=%23RISC-V%E8%99%9A%E6%8B%9F%E5%8C%96%E5%AE%9E%E6%88%98%23" data-hide=""><span class="surl-text">#RISC-V虚拟化实战#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%9F%BA%E4%BA%8ERISC-V%E6%89%8B%E6%92%B8Hypervisor%23&amp;extparam=%23%E5%9F%BA%E4%BA%8ERISC-V%E6%89%8B%E6%92%B8Hypervisor%23" data-hide=""><span class="surl-text">#基于RISC-V手撸Hypervisor#</span></a><br><br>如何基于RISC-V H-extension，手撸一个Hypervisor，并成功运行Linux？<br><br>这位开发者在自研操作系统Starina实现虚拟化支持，并记录了全过程。<br><br>整个过程从零开始，一步步搭起一个能跑Linux的轻量虚拟机环境，架构类似WSL2，但运行在完全自研的系统内核上。<br><br>Step 1：进入Guest模式  <br>RISC-V通过H-extension提供VS-mode（Virtual Supervisor），相当于Guest OS的内核态。设置好CSR（特别是 `hstatus.SPV`）并执行 `sret`，CPU就能跳入Guest态，成功触发了instruction guest-page fault，说明模式切换成功。<br><br>Step 2：执行第一条ecall  <br>为了验证Guest能正常运行指令，加载了一个简单程序跑 `ecall`。这需要先构建Guest页表（使用Sv39x4等模式），并配置 `hgatp` 映射地址。Guest trap成功说明访问机制生效。<br><br>Step 3：Hello World from Guest  <br>用RISC-V汇编写了一个最小程序，通过SBI（类似BIOS接口）输出字符“Hi!”。这个测试验证了系统调用处理、字符输出和中断处理链都通了。<br><br>Step 4：尝试跑Linux内核  <br>构建了RISC-V Linux Image，复制进Guest内存，CPU跳转后成功执行，但马上遇到空指针异常，crash在 `__pi_fdt32_ld` —— 问题出在没有提供Device Tree（设备树）。<br><br>Step 5：补上Device Tree  <br>用RustVMM的 `vm-fdt` crate 构建最小设备树，指定内存、CPU等信息，让Linux知道自己运行在什么样的“硬件”上。这个crate支持 `no_std`，直接集成到内核非常方便。<br><br>Step 6：支持rdtime读取  <br>Linux读 `rdtime` 报错，原因是 `hcounteren` CSR没开启。补上之后，Guest内核能正常读取时间戳。这也体现了RISC-V对寄存器权限的严格控制。<br><br>Step 7：支持定时器与中断注入  <br>Linux内核启动时需要确认系统时钟在“跳动”。通过实现 `sstc` 扩展 + 设置 `hideleg`，Guest能正确接收到timer中断。第一次成功向Guest注入中断，超关键。<br><br>Step 8：模拟MMIO设备  <br>想要真正跑起来，还得让Linux能看到设备。用MMIO机制模拟了常见硬件（如PLIC中断控制器、virtio-blk磁盘、virtio-net网卡），Guest访问设备地址时触发trap，Hypervisor拦截并模拟读写。<br><br>Step 9：接入virtio-fs文件系统  <br>为了更好地和Starina集成，选用了 `virtio-fs` 替代传统的 `virtio-blk`，通过FUSE协议实现Host与Guest共享文件系统，为后续文件交换打下基础。<br><br>九大步在RISC-V上写Hypervisor，Linux就能跑在你自己写的OS里。<br><br>全程都能在QEMU中完成，QEMU已经支持RISC-V H-extension模拟，再加上GDB调试支持，连Hypervisor和Guest的栈回溯都一目了然，开发体验非常丝滑。<br><br>原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fseiya.me%2Fblog%2Friscv-hypervisor" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1iwaswm9hj30nc0oa12s.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i1iwazbo3jj30mj0nkn5y.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

该开发者基于RISC-V H-extension在自研操作系统Starina上实现了Hypervisor虚拟化支持，成功运行Linux。过程包括：1）通过H-extension进入Guest模式；2）构建Guest页表并处理系统调用；3）实现基础SBI输出；4）加载Linux内核并解决设备树缺失问题；5）支持时间读取和定时器中断；6）模拟MMIO设备和virtio-fs文件系统。最终在QEMU模拟环境中完成了从零构建的轻量级虚拟机环境，架构类似WSL2但完全自研。开发过程充分利用了RISC-V的虚拟化扩展和QEMU的调试支持。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-17T22:02:27Z
- **目录日期**: 2025-05-17
