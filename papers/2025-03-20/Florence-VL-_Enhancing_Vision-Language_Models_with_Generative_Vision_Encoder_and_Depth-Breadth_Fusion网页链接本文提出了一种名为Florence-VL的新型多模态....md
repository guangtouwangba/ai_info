# Florence-VL: Enhancing Vision-Language Models with Generative Vision Encoder and Depth-Breadth Fusion网页链接本文提出了一种名为Florence-VL的新型多模态...

**URL**: https://weibo.com/1870858943/P4uA6ySTO

## 原始摘要

Florence-VL: Enhancing Vision-Language Models with Generative Vision Encoder and Depth-Breadth Fusion<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.aminer.cn%2Fpub%2F67526da4ae8580e7ff3d4f67%2F%3Ff%3Dwb" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>本文提出了一种名为Florence-VL的新型多模态大型语言模型（MLLM），通过集成生成性视觉基础模型Florence-2来增强视觉表征。与常见的CLIP风格的视觉变换器不同，Florence-2能够捕捉到视觉特征的不同层次和方面，使其能够适应多种下游任务。研究团队设计了一种新颖的特征融合架构和创新的训练方法，将Florence-2的视觉特征有效融合到预训练的语言模型Phi 3.5和LLama 3中。特别提出了一种“深度-广度融合（DBFusion）”方法，用于融合不同深度和多个提示下提取的视觉特征。模型的训练包括整个模型的端到端预训练，以及对投影层和语言模型的微调，使用精心设计的包含高质量图像标题和指令调整对的多种开源数据集。定量分析和视觉化结果显示，Florence-VL在视觉语言对齐方面优于流行的视觉编码器，其中增强的深度和广度发挥了重要作用。Florence-VL在各种多模态和视觉为中心的基准测试中，如通用VQA、感知、幻觉、OCR、图表、知识密集型理解等方面，都取得了显著改进。为促进未来研究，研究团队开源了模型和完整的训练方法。<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%A8%A1%E5%9E%8B%E8%AE%AD%E7%BB%83%23&amp;extparam=%23%E6%A8%A1%E5%9E%8B%E8%AE%AD%E7%BB%83%23" data-hide=""><span class="surl-text">#模型训练#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%A4%A7%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#大语言模型#</span></a><a href="https://m.weibo.cn/p/index?extparam=%E4%BA%BA%E5%B7%A5%E6%99%BA%E8%83%BD&amp;containerid=100808f068f0dad74789bee210163c40a4b50d" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">人工智能</span></a><a href="https://m.weibo.cn/p/index?extparam=%E7%A1%95%E5%A3%AB%E8%AE%BA%E6%96%87&amp;containerid=1008084cacf38f5903dc7b04550404d0bd3608" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">硕士论文</span></a><a href="https://m.weibo.cn/p/index?extparam=%E8%AE%BA%E6%96%87%E5%86%99%E4%BD%9C&amp;containerid=1008084f70c9f305ba97c50dbca8c25c8747d7" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">论文写作</span></a><img style="" src="https://tvax3.sinaimg.cn/large/6f830abfly1hwhxwdzsabj22cf17zqv5.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

本文介绍了Florence-VL，一种新型多模态大型语言模型（MLLM），通过集成生成性视觉基础模型Florence-2来增强视觉表征。Florence-2能够捕捉视觉特征的不同层次和方面，适应多种下游任务。研究团队设计了特征融合架构和训练方法，将Florence-2的视觉特征融合到预训练的语言模型Phi 3.5和LLama 3中，提出“深度-广度融合（DBFusion）”方法。Florence-VL在视觉语言对齐方面优于流行的视觉编码器，并在多模态和视觉为中心的基准测试中取得显著改进。模型和训练方法已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T12:06:26Z
- **目录日期**: 2025-03-20
