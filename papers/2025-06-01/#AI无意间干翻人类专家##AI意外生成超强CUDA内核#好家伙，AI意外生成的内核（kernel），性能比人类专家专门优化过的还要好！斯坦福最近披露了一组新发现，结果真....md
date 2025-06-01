# #AI无意间干翻人类专家##AI意外生成超强CUDA内核#好家伙，AI意外生成的内核（kernel），性能比人类专家专门优化过的还要好！斯坦福最近披露了一组新发现，结果真...

**URL**: https://weibo.com/6105753431/PumTnEgOB

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E6%97%A0%E6%84%8F%E9%97%B4%E5%B9%B2%E7%BF%BB%E4%BA%BA%E7%B1%BB%E4%B8%93%E5%AE%B6%23&amp;extparam=%23AI%E6%97%A0%E6%84%8F%E9%97%B4%E5%B9%B2%E7%BF%BB%E4%BA%BA%E7%B1%BB%E4%B8%93%E5%AE%B6%23" data-hide=""><span class="surl-text">#AI无意间干翻人类专家#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E6%84%8F%E5%A4%96%E7%94%9F%E6%88%90%E8%B6%85%E5%BC%BACUDA%E5%86%85%E6%A0%B8%23&amp;extparam=%23AI%E6%84%8F%E5%A4%96%E7%94%9F%E6%88%90%E8%B6%85%E5%BC%BACUDA%E5%86%85%E6%A0%B8%23" data-hide=""><span class="surl-text">#AI意外生成超强CUDA内核#</span></a><br><br>好家伙，AI意外生成的内核（kernel），性能比人类专家专门优化过的还要好！<br><br>斯坦福最近披露了一组新发现，结果真的太亮眼了。【图1】<br><br>由AI优化的内核，在常见深度学习操作上，翻倍超越原生PyTorch，性能至多可以提升近400%——<br><br>- 矩阵乘法（Matmul，FP32）：性能达到PyTorch&nbsp;torch.matmul的101.3%。<br><br>- 二维卷积（Conv2D）：性能达到&nbsp;torch.nn.Conv2D的179.9%。<br><br>- Softmax：性能达到&nbsp;torch.softmax的111.8%。<br><br>- 层归一化（LayerNorm）：性能达到torch.nn.LayerNorm的484.4%。<br><br>- Conv2D+ReLU+MaxPool组合操作：性能达到PyTorch参考实现的290.1%，以及torch.compile()参考实现的189.0%。<br><br>更惊人的是，这一切都是意外实现的。<br><br>研究团队本来的目标是生成合成数据以训练内核生成模型。<br><br>结果发现，仅在测试阶段生成的合成数据本身，竟然可以生成性能非常优秀的内核。【图2】<br><br>围观网友：没想到AI也要取代内核工程师了。【图3】<br><br>还有人发现，除了性能大幅提升外，研究团队采用的方法也非常有趣：<br><br>他们没有简单的在操作上逐步优化（类似于爬坡算法），而是在每次迭代之间加入了一个语言推理的步骤，通过这种方式鼓励搜索过程更加多样化。<br><br>也就是说，他们是让系统在每次改进时通过类似“思考”的方式产生更多想法，从而找到更好的解决方案。【图4】<br><br>具体如何实现，一起来看：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FLFZhiacSqkaTkuMHWdzE1A" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">斯坦福意外用AI生成超强CUDA内核，性能比人类专家优化得还要好！翻倍碾压原生PyTorch，华人主创</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i1yogxui4dj30zk0jzgqx.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1yohe4dgpj30ws0a2ae7.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1yohgbjd7j30wo0swwrm.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1yokctslbj30jq06ydim.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

斯坦福研究团队意外发现，AI生成的CUDA内核性能超越人类专家优化版本，在多项深度学习操作中表现亮眼：矩阵乘法达PyTorch的101.3%，二维卷积达179.9%，层归一化更达到484.4%。该突破源于测试阶段生成的合成数据意外具备优秀性能。研究采用创新方法，在迭代间加入语言推理步骤以增强搜索多样性，类似"思考"过程帮助找到更优解。这一发现可能颠覆传统内核优化模式，引发对AI取代内核工程师的讨论。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-01T14:04:51Z
- **目录日期**: 2025-06-01
