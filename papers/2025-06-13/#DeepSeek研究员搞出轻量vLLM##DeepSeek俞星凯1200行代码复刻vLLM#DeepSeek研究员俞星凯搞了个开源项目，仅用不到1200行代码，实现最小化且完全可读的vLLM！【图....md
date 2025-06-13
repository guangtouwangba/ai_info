# #DeepSeek研究员搞出轻量vLLM##DeepSeek俞星凯1200行代码复刻vLLM#DeepSeek研究员俞星凯搞了个开源项目，仅用不到1200行代码，实现最小化且完全可读的vLLM！【图...

**URL**: https://weibo.com/6105753431/PwmcfiVop

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23DeepSeek%E7%A0%94%E7%A9%B6%E5%91%98%E6%90%9E%E5%87%BA%E8%BD%BB%E9%87%8FvLLM%23&amp;extparam=%23DeepSeek%E7%A0%94%E7%A9%B6%E5%91%98%E6%90%9E%E5%87%BA%E8%BD%BB%E9%87%8FvLLM%23" data-hide=""><span class="surl-text">#DeepSeek研究员搞出轻量vLLM#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23DeepSeek%E4%BF%9E%E6%98%9F%E5%87%AF1200%E8%A1%8C%E4%BB%A3%E7%A0%81%E5%A4%8D%E5%88%BBvLLM%23&amp;extparam=%23DeepSeek%E4%BF%9E%E6%98%9F%E5%87%AF1200%E8%A1%8C%E4%BB%A3%E7%A0%81%E5%A4%8D%E5%88%BBvLLM%23" data-hide=""><span class="surl-text">#DeepSeek俞星凯1200行代码复刻vLLM#</span></a><br><br>DeepSeek研究员俞星凯搞了个开源项目，仅用不到1200行代码，实现最小化且完全可读的vLLM！【图1】<br><br>项目名为Nano-vLLM（纳米级-vLLM），有三大特点：<br><br>- 快速离线推理：推理速度可与vLLM相媲美<br>    <br>- 可读性强的代码库：基于不到1200行Python代码实现，简洁干净<br>    <br>- 优化套件：包含前缀缓存、Torch compilation 、CUDA graph等<br><br>下面是vLLM与Nano-vLLM在不同硬件和模型配置下的基准测试情况：【图2】vLLM略微领先。<br><br>该测试在RTX 4070硬件、Qwen3-0.6B模型环境中，设置了256个序列的总请求数，输入和输出长度均在100-1024个 token间随机采样。<br><br>二者输出token量相同，vLLM耗时98.95秒、吞吐量为1353.86 tokens/s，Nano-vLLM耗时101.90秒、吞吐量1314.65tokens/s。<br><br>DeepSeek研究员俞星凯，他2021年获得南京大学计算机科学与技术系学士学位，同年又被南京大学免试录取为硕士研究生，在校他同时也是由周志华教授领导的LAMDA团队的成员。<br><br>详情点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FZ8lYMahi-G7sq9t7b9W6kg" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">DeepSeek研究员1200行代码复刻vLLM，H800硬件实测性能反超原版</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2drljic38j30zk09pwhw.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2drlq1lhcj30ui0hugqn.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2drpc3wvvj30u00f2jyz.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

DeepSeek研究员俞星凯开源了轻量级项目Nano-vLLM，仅用1200行Python代码复刻了vLLM的核心功能。该项目具有高效推理（吞吐量1314.65 tokens/s，接近原版vLLM的1353.86）、代码简洁可读的特点，并整合了前缀缓存、Torch编译等优化技术。测试显示，在RTX 4070硬件上运行Qwen3-0.6B模型时，两者性能差距仅2.3%。俞星凯系南京大学LAMDA团队成员，该项目实现了vLLM的高效精简复现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T09:03:26Z
- **目录日期**: 2025-06-13
