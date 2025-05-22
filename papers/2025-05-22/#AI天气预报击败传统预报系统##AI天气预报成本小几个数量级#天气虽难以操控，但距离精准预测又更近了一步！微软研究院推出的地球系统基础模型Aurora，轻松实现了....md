# #AI天气预报击败传统预报系统##AI天气预报成本小几个数量级#天气虽难以操控，但距离精准预测又更近了一步！微软研究院推出的地球系统基础模型Aurora，轻松实现了...

**URL**: https://weibo.com/6105753431/Pt0wn3nBT

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%A4%A9%E6%B0%94%E9%A2%84%E6%8A%A5%E5%87%BB%E8%B4%A5%E4%BC%A0%E7%BB%9F%E9%A2%84%E6%8A%A5%E7%B3%BB%E7%BB%9F%23&amp;extparam=%23AI%E5%A4%A9%E6%B0%94%E9%A2%84%E6%8A%A5%E5%87%BB%E8%B4%A5%E4%BC%A0%E7%BB%9F%E9%A2%84%E6%8A%A5%E7%B3%BB%E7%BB%9F%23" data-hide=""><span class="surl-text">#AI天气预报击败传统预报系统#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%A4%A9%E6%B0%94%E9%A2%84%E6%8A%A5%E6%88%90%E6%9C%AC%E5%B0%8F%E5%87%A0%E4%B8%AA%E6%95%B0%E9%87%8F%E7%BA%A7%23&amp;extparam=%23AI%E5%A4%A9%E6%B0%94%E9%A2%84%E6%8A%A5%E6%88%90%E6%9C%AC%E5%B0%8F%E5%87%A0%E4%B8%AA%E6%95%B0%E9%87%8F%E7%BA%A7%23" data-hide=""><span class="surl-text">#AI天气预报成本小几个数量级#</span></a><br><br>天气虽难以操控，但距离精准预测又更近了一步！<br><br>微软研究院推出的地球系统基础模型Aurora，轻松实现了几个关键预测领域的SoTA，计算成本还要小好几个数量级：<br><br>- 0.4° 分辨率的 5 天全球空气污染预报：在74%的预测指标上优于高资源消耗的大气化学数值模拟；【图2】<br>- 0.25° 分辨率的10天全球海浪预报：在86%的目标上优于成本高昂的数值模型。【图3】<br>- 5天热带气旋路径预报：在100%的预测指标上优于七大业务预报中心【图4】<br>- 0.1° 分辨率的10天全球天气预报：在92%的预测指标上领先最先进数值模型，且极端天气预测性能显著提升【图5】<br><br>可以看出，多领域预测能力以及高分辨率预测是Aurora的强项。<br><br>那么，功能如此强悍的Aurora，具体具有什么样的结构呢？<br><br>该模型由三部分组成：<br><br>（1）编码器：将不同类型的数据转换为统一的三维潜空间表征；<br>（2）处理器：让这些表征按照时间顺序自然地发展和变化；<br>（3）解码器：将处理后的表征重新转换成我们能看懂的实际预测结果。<br><br>其中，处理器采用了3D Swin Transformer U-Net架构，编码器与解码器则基于3D Perceiver模块构建。【图6】<br><br>整个训练结果可以分为两个阶段来看：<br><br>1、预训练阶段：使用超过一百万小时的多样化地球系统数据进行预训练，通过增加预训练数据量和模型规模，模型性能还有望进一步提升。<br><br>2、微调阶段：完成预训练后，模型可利用习得的通用表征，通过第二阶段训练高效适配新任务、新数据集及新变量。<br><br>值得一提的是，Aurora可以以极低的成本微调，适配多种地球系统预测任务，有望用于海洋环流预测、植被生长监测、极端天气预警等领域。<br><br>论文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2405.13063" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1oc2pa9mgj30zk0hbjxh.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1oc2rrns7j30zk0oxqlg.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1oc2tspiej30ze0sy7pr.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1oc2yidoyj30zk0iadmq.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1oc30dhdrj30zk0of4dl.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1oc33otdej30zk0ratm4.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

微软研究院推出的Aurora地球系统基础模型在天气预报领域取得突破性进展。该AI模型在多个关键预测任务中超越传统数值模型：5天全球空气污染预报（74%指标更优）、10天海浪预报（86%更优）、热带气旋路径预报（100%领先）及10天全球天气预报（92%指标领先）。Aurora采用3D Swin Transformer U-Net架构，通过百万小时数据预训练和低成本微调实现高性能。其优势包括超高分辨率预测（最低0.1°）、极端天气预警能力，且计算成本比传统方法低数个数量级，可扩展应用于海洋、生态等多领域预测。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T20:03:55Z
- **目录日期**: 2025-05-22
