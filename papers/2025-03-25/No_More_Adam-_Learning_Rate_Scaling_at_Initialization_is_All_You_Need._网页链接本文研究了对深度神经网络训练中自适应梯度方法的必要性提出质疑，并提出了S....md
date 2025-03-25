# No More Adam: Learning Rate Scaling at Initialization is All You Need. 网页链接本文研究了对深度神经网络训练中自适应梯度方法的必要性提出质疑，并提出了S...

**URL**: https://weibo.com/1870858943/Pk06MwyNe

## 原始摘要

No More Adam: Learning Rate Scaling at Initialization is All You Need. <a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.aminer.cn%2Fpub%2F6762315eae8580e7ff8ed69e%2Fno-more-adam-learning-rate-scaling-at-initialization-is-all-you-need" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>本文研究了对深度神经网络训练中自适应梯度方法的必要性提出质疑，并提出了SGD-SaI方法。该方法是对带动量的随机梯度下降（SGDM）的简单而有效的改进。SGD-SaI通过对不同参数组在初始化时进行学习率调整，依据各自的梯度信号噪声比（g-SNR）来指导。这种方法无需依赖自适应二阶动量，就能从训练的第一步起防止训练不平衡，并将优化器的内存使用量减半，相比AdamW有显著优势。SGD-SaI不仅在多种基于Transformer的任务中与AdamW表现相当或更佳，解决了使用SGD训练Transformer的难题，还在ImageNet-1K分类任务和大型语言模型（LLM）的GPT-2预训练中表现出色，具备对超参数变化的鲁棒性和实际应用的广泛性。此外，它在LoRA微调等任务中的表现也优于现有最优优化器。在内存效率方面，SGD-SaI在完整精度训练设置下，相比AdamW为GPT-2节省了5.93GB内存，为Llama2-7B节省了25.15GB内存。<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#大模型#</span></a><a href="https://m.weibo.cn/p/index?extparam=%E4%BA%BA%E5%B7%A5%E6%99%BA%E8%83%BD&amp;containerid=100808f068f0dad74789bee210163c40a4b50d" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">人工智能</span></a><img style="" src="https://tvax4.sinaimg.cn/large/6f830abfly1hzrut5jebij21na0z04qp.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

该论文提出SGD-SaI方法，质疑深度神经网络训练中自适应梯度方法(如Adam)的必要性。该方法通过初始化时基于梯度信噪比(g-SNR)调整各参数组学习率，改进带动量的SGD(SGDM)，无需自适应二阶动量即可避免训练不平衡，并将内存使用减半。实验显示，SGD-SaI在Transformer任务、ImageNet分类、GPT-2预训练和LoRA微调中表现优于或持平AdamW，尤其在大模型训练中显著节省内存(GPT-2节省5.93GB，Llama2-7B节省25.15GB)，同时保持超参数鲁棒性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-25T17:04:31Z
- **目录日期**: 2025-03-25
