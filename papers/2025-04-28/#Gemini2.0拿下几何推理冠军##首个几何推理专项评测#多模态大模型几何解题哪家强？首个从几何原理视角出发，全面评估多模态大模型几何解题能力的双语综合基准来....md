# #Gemini2.0拿下几何推理冠军##首个几何推理专项评测#多模态大模型几何解题哪家强？首个从几何原理视角出发，全面评估多模态大模型几何解题能力的双语综合基准来...

**URL**: https://weibo.com/6105753431/PpmPgnIWD

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Gemini2.0%E6%8B%BF%E4%B8%8B%E5%87%A0%E4%BD%95%E6%8E%A8%E7%90%86%E5%86%A0%E5%86%9B%23&amp;extparam=%23Gemini2.0%E6%8B%BF%E4%B8%8B%E5%87%A0%E4%BD%95%E6%8E%A8%E7%90%86%E5%86%A0%E5%86%9B%23" data-hide=""><span class="surl-text">#Gemini2.0拿下几何推理冠军#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%A6%96%E4%B8%AA%E5%87%A0%E4%BD%95%E6%8E%A8%E7%90%86%E4%B8%93%E9%A1%B9%E8%AF%84%E6%B5%8B%23&amp;extparam=%23%E9%A6%96%E4%B8%AA%E5%87%A0%E4%BD%95%E6%8E%A8%E7%90%86%E4%B8%93%E9%A1%B9%E8%AF%84%E6%B5%8B%23" data-hide=""><span class="surl-text">#首个几何推理专项评测#</span></a><br><br>多模态大模型几何解题哪家强？<br><br>首个从几何原理视角出发，全面评估多模态大模型几何解题能力的双语综合基准来了！<br><br>GeoSense，系统评测多模态大模型在几何原理识别和应用中的表现，评测基准的数据和评测代码均已开源。<br><br>其背后团队来自淘天集团算法技术-未来生活实验室团队。、<br><br>评测系统基于五层知识架构，涵盖148个几何原理，覆盖平面与立体几何。数据集中包含1789道题，中英双语精细标注，共记录5556个几何原理的应用细节，由23位几何专业研究生参与标注与审核。【图1】<br><br>GeoSense采用创新评估体系：【图2】<br><br>- GPI（几何原理识别得分）<br>- GPA（几何原理应用得分）<br>- ACC（最终答案正确率）<br><br>该团队对多个开源和闭源模型进行了全面评测和分析，并给出了这些模型在GPI、GPA和ACC三个指标的平均值上的排名情况。<br><br>Gemini-2.0-Pro-Flash在评测中表现最佳，开源阵营中Qwen-VL系列也表现亮眼。<br><br>推理增强模型如QVQ-72B展现潜力，但复杂问题中准确率仍受限。平面几何理解则是所有模型的共同短板，比如Claude3.5在二维图形理解上识别和应用准确率都偏低。<br><br>另一个趋势是，模型规模普遍越大，推理表现越优，如Qwen2.5-VL系列从7B到72B的指标大幅提升。<br><br>分析发现，GPI和GPA的下降都会直接拖累ACC。且复杂问题下，模型在几何原理识别环节掉链子更严重，导致整体推理能力下降。<br><br>从知识类型来看，模型在公式类问题上表现更好，但在涉及定义和定理的理解时仍有明显不足。<br><br>特别在平面几何领域，识别混淆性强的原理如相似三角形，全等三角形等成为大难题，极大限制了解题能力。<br><br>最后，GeoSense团队揭示了当前多模态大模型在几何推理上的短板，也指明了优化方向——提升原理识别能力，才能真正推动推理能力进化。原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FPz83EBkBvXfY8tlp2iV98w" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0wnl4zgmjj30lc0jawkx.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0wnl67hkjj30zk0hxh30.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

GeoSense是首个针对多模态大模型几何推理能力的双语评测基准，由淘天集团团队开发。该基准基于五层知识架构，涵盖148个几何原理和1789道中英双语题目。评测显示，Gemini-2.0-Pro-Flash表现最佳，开源模型中Qwen-VL系列突出。研究发现模型规模越大推理能力越强，但平面几何仍是共同短板，尤其在相似/全等三角形等原理识别上存在困难。评测指出提升原理识别能力是优化几何推理的关键。数据和代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T22:03:24Z
- **目录日期**: 2025-04-28
