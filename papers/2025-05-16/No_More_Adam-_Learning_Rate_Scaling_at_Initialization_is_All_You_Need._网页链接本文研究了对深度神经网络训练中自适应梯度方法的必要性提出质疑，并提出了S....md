# No More Adam: Learning Rate Scaling at Initialization is All You Need. 网页链接本文研究了对深度神经网络训练中自适应梯度方法的必要性提出质疑，并提出了S...

**URL**: https://weibo.com/1870858943/Pk06MwyNe

## 原始摘要

No More Adam: Learning Rate Scaling at Initialization is All You Need. <a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.aminer.cn%2Fpub%2F6762315eae8580e7ff8ed69e%2Fno-more-adam-learning-rate-scaling-at-initialization-is-all-you-need" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>本文研究了对深度神经网络训练中自适应梯度方法的必要性提出质疑，并提出了SGD-SaI方法。该方法是对带动量的随机梯度下降（SGDM）的简单而有效的改进。SGD-SaI通过对不同参数组在初始化时进行学习率调整，依据各自的梯度信号噪声比（g-SNR）来指导。这种方法无需依赖自适应二阶动量，就能从训练的第一步起防止训练不平衡，并将优化器的内存使用量减半，相比AdamW有显著优势。SGD-SaI不仅在多种基于Transformer的任务中与AdamW表现相当或更佳，解决了使用SGD训练Transformer的难题，还在ImageNet-1K分类任务和大型语言模型（LLM）的GPT-2预训练中表现出色，具备对超参数变化的鲁棒性和实际应用的广泛性。此外，它在LoRA微调等任务中的表现也优于现有最优优化器。在内存效率方面，SGD-SaI在完整精度训练设置下，相比AdamW为GPT-2节省了5.93GB内存，为Llama2-7B节省了25.15GB内存。<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#大模型#</span></a><a href="https://m.weibo.cn/p/index?extparam=%E4%BA%BA%E5%B7%A5%E6%99%BA%E8%83%BD&amp;containerid=100808f068f0dad74789bee210163c40a4b50d" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">人工智能</span></a><img style="" src="https://tvax4.sinaimg.cn/large/6f830abfly1hzrut5jebij21na0z04qp.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

这篇论文提出了一种名为SGD-SaI的新优化方法，挑战了深度神经网络训练中必须使用自适应梯度方法（如Adam）的传统观念。SGD-SaI基于带动量的随机梯度下降（SGDM），通过在初始化阶段根据不同参数组的梯度信号噪声比（g-SNR）调整学习率，避免了训练不平衡问题。相比AdamW，该方法内存使用减半，在Transformer、ImageNet分类、GPT-2预训练等任务中表现相当或更好，同时显著节省内存（如为GPT-2节省5.93GB，Llama2-7B节省25.15GB）。该方法还具有超参数鲁棒性，在LoRA微调等任务中优于现有优化器。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T14:05:19Z
- **目录日期**: 2025-05-16
